import asyncio
import boto3
import graphene
import logging

from typing import List

from pyramid.threadlocal import get_current_request

from graphene import resolve_only_args

import youtube_dl

from youtube_dl.utils import (
    DEFAULT_OUTTMPL
)

from ..utils import to_encoded_str

log = logging.getLogger(__name__)

BUCKET_NAME = 'kamestery-audio-bucket'

AUDIO_POSTPROCESSOR = dict(
    key='FFmpegExtractAudio',
    preferredcodec='mp3',
    preferredquality='192'
)

def get_ydl_opts(**opts):
    ydl_opts = dict(
        format='bestaudio/best',
        postprocessors=[AUDIO_POSTPROCESSOR],
        verbose=True,
        progress_hooks=[],
        outtmpl='/tmp/%(title)s-%(id)s.%(ext)s'
    )
    ydl_opts.update(opts)
    return ydl_opts

async def get_media_info(ydl_opts: dict, media_urls: List[str]):
    media_info = []
    ydl_opts = get_ydl_opts(**dict(logger=log))
    ydl = youtube_dl.YoutubeDL(ydl_opts)
    for url in media_urls:
        with ydl:
            info = ydl.extract_info(
                url, download=False  # We just want to extract the info
            )
            media_info.append(info)
    return media_info

def download_media(ydl_opts: dict, media_urls: List[str]):
    with youtube_dl.YoutubeDL(ydl_opts) as ydl:
        ydl.download(media_urls)

async def s3upload(filename: str, file_key: str, file_path: str, **opts):
    # Create an S3 client
    s3 = boto3.client('s3')
    options = {
        'ACL': 'public-read'
    }
    options.update(**opts)
    s3.upload_file(file_path, BUCKET_NAME, filename, ExtraArgs=options)
    s3_http_url = s3.generate_presigned_url(
        ClientMethod='get_object',
        Params={
            'Bucket': BUCKET_NAME,
            'Key': file_key
        }
    )
    return s3_http_url

class MediaQuery():
    media_query = graphene.Field(
        graphene.List(graphene.String),
        media_urls=graphene.List(
            graphene.NonNull(graphene.String),
            description='Media URLs'
        ),
        format = graphene.String(description='Format', default_value='bestaudio/best'),
        output_directory = graphene.String(description='Output Directory', default_value='/tmp')
    )

    @resolve_only_args
    async def resolve_media_query(self, media_urls: List[str], format, output_directory):
        log.debug(f'Downloading {media_urls} and extracting {format} into directory {output_directory}')
        ydl_opts = get_ydl_opts(**dict(format=format, outtmpl=f'{output_directory}/{DEFAULT_OUTTMPL}', logger=log))
        # download_media(ydl_opts, media_urls) # Download Asynchronously
        request = get_current_request()
        audio_files = []
        download_urls = []
        upload_tasks = []
        media_info = await get_media_info(ydl_opts, media_urls)
        for info in media_info:
            output_file = f'{output_directory}/{info["title"]}-{info["id"]}.{info["ext"]}'
            log.debug(f'Video File: {output_file}')
            audio_file = f'{output_directory}/{info["title"]}-{info["id"]}.mp3'
            log.debug(f'Audio File: {audio_file}')
            audio_files.append(audio_file)
            download_url = request.route_url('media', media_id=to_encoded_str(audio_file))
            download_urls.append(download_url)
            audio_file_key = info["id"]
            audio_filename = f'{info["title"]}-{audio_file_key}.mp3'.replace(' ', '-')
            upload_tasks.append(
                s3upload(audio_filename, audio_file_key, audio_file)
            )

        asyncio.get_event_loop().run_until_complete(asyncio.wait(
            upload_tasks
        ))

        return download_urls
