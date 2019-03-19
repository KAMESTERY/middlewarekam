package utils

import (
	"bytes"
	"html/template"
	"os"
	"regexp"
	"strings"
)

const (
	SHORTCODE_REGEX = "\\[\\[([a-z-]+)(.*?)\\]\\]" // In the form [[command param param param]]
	SHORTCODE_ERROR = "[[ERROR]]"
)

var (
	utilsshortcodes_logger = NewLogger("utilsshortcodes")
	re_sht = regexp.MustCompile(SHORTCODE_REGEX)
)

func shortcodes_tpl() *template.Template {
	return template.Must(template.New("").Funcs(template.FuncMap{"include": include, "env": env}).Parse(SHORTCODES))
}

func include(name string, data interface{}) (ret template.HTML, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = shortcodes_tpl().ExecuteTemplate(buf, name, data)
	if err != nil {
		utilsshortcodes_logger.Errorf("ERROR:::: %+v", err.Error())
	}
	ret = template.HTML(buf.String())
	return
}

func env(envVarName string) (envVarValue string) {
	if envVarValue = os.Getenv(envVarName); len(envVarValue) == 0 {
		envVarValue = "No Found"
	}
	return
}

func renderShtCode(tmpl string, data interface{}) (writtenString string) {
	buf := new(bytes.Buffer)
	err := shortcodes_tpl().ExecuteTemplate(buf, tmpl, data)
	if err != nil {
		utilsshortcodes_logger.Errorf("SHORTCODE_ERROR:::: %+v", err)
		writtenString = SHORTCODE_ERROR
		return
	}
	writtenString = buf.String()
	return
}

func ExpandShortcodes(document string) string {

	matches := re_sht.FindAllString(document, -1)

	for _, match := range matches {
		shtcd_params := strings.Replace(match, "[[", "", -1)
		shtcd_params = strings.Replace(shtcd_params, "]]", "", -1)
		params := strings.Split(shtcd_params, " ")
		data := struct {
			Tmpl string
			Args []string
		}{params[0], params[1:]}
		chunk := renderShtCode("shortcode", data)
		document = strings.Replace(document, match, chunk, -1)
	}

	return document
}

const SHORTCODES = `
{{define "shortcode"}}
    {{include .Tmpl .}}
{{ end }}
{{define "amazoncontextwidget"}}
<script type="text/javascript" language="javascript">
    amzn_assoc_ad_type = "contextual";
    amzn_assoc_tracking_id = "{{env "AMZN_ASSOC_TRACKING_ID"}}";
    amzn_assoc_marketplace = "amazon";
    amzn_assoc_region = "US";
    amzn_assoc_placement = "{{env "AMZN_ASSOC_PLACEMENT"}}";
    amzn_assoc_linkid = "{{env "AMZN_ASSOC_LINKID"}}";
    amzn_assoc_emphasize_categories = "";
    amzn_assoc_fallback_products = "";
    amzn_assoc_width = "300";
    amzn_assoc_height = "250";
</script>
<script type="text/javascript" language="javascript" src="//z-na.amazon-adsystem.com/widgets/q?ServiceVersion=20070822&Operation=GetScript&ID=OneJS&WS=1&MarketPlace=US&source=ac"></script>
{{ end }}
{{define "youtube"}}
<div class='embed-container' style="padding-bottom: 56.25%; /* 16:9 */ padding-top: 25px; width: 480; height: 385;">
    <iframe src='//www.youtube.com/embed/{{ index .Args 0 }}' frameborder='0' allowfullscreen></iframe>
</div>
{{ end }}
{{define "pinterest"}}
<a data-pin-do="embedPin" href="//www.pinterest.com/pin/{{ index .Args 0 }}/"></a>
{{ end }}
{{define "itunes"}}
<div class="embed-container">
    <iframe
            src="//widgets.itunes.apple.com/widget.html?c=us&brc=FFFFFF&blc=FFFFFF&trc=FFFFFF&tlc=FFFFFF&d=&t=&m=music&e=album&w=250&h=300&ids={{ index .Args 0 }}&wt=discovery&partnerId=&affiliate_id={{env "APPLE_AFF_ID"}}&at=1000l9ug&ct="
            frameborder=0
            style="overflow-x:hidden;overflow-y:hidden;width:250px;height: 300px;border:0px"></iframe>
</div>
{{ end }}
{{define "soundcloud"}}
<div class="embed-container">
    <iframe width="100%" height="450" scrolling="no" frameborder="no"
            src="//w.soundcloud.com/player/?url=https%3A//api.soundcloud.com/tracks/{{ index .Args 0 }}&amp;auto_play=false&amp;hide_related=false&amp;show_comments=true&amp;show_user=true&amp;show_reposts=false&amp;visual=true"></iframe>
</div>
{{ end }}
{{define "instagram"}}
<div class='embed-container'>
    <iframe width="100%" height="100%" src='//instagram.com/p/{{ index .Args 0 }}/embed/' frameborder='0' scrolling='no' allowtransparency='true'></iframe>
</div>
{{ end }}
{{define "dailymotion"}}
<div class='embed-container' style="padding-bottom: 56.25%; /* 16:9 */ padding-top: 25px;">
    <iframe src='//www.dailymotion.com/embed/video/{{ index .Args 0 }}' frameborder='0' webkitAllowFullScreen mozallowfullscreen allowFullScreen></iframe>
</div>
{{ end }}
{{define "vimeo"}}
<div class='embed-container' style="padding-bottom: 56.25%; /* 16:9 */ padding-top: 25px;">
    <iframe src='//player.vimeo.com/video/{{ index .Args 0 }}' frameborder='0' webkitAllowFullScreen mozallowfullscreen allowFullScreen></iframe>
</div>
{{ end }}
{{define "link"}}
<a href="/{{ index .Args 0 }}/{{ index .Args 1 }}">{{ index .Args 2 }}</a>
{{ end }}
`
