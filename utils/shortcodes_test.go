package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// https://regex-golang.appspot.com/assets/html/index.html
// http://www.regexplanet.com/advanced/golang/index.html

func TestExpandShortcodes(t *testing.T) {
	expanded := ExpandShortcodes(test_document)
	fmt.Println(expanded)
	assert.True(t, len(expanded) > 0)
	assert.True(t, !strings.Contains(expanded, SHORTCODE_ERROR))
}

const test_document = `
L'Afrique par essence est basée sur la communauté, ou chacun se soucie de chacun.
Les Africains sont chaleureux, hospitaliers et accueillants.
Les habitations sont construites de telles sortes que bien que l'individu a un certain espace prive, il ne soit pas coupe des autres.
Il y a une fameux proverbe africain qui dit que "il faut un village pour élever un enfant" parce que chez les Africains, l'enfant appartient a toute la société mais pas aux parents seulement. Ceci est possible qu'a cause de l'esprit de communauté. Une chose particulièrement frappante c'est que pour un étranger non africain il ne saura dire quel enfant appartient a quel parent, et qui sont les résidents de telle ou telle maison. Les voisins se rendent visite, mangent ensemble.
Tout est occasion de celebration et de réjouissance; les femmes chantent et dansent pendant qu'elles préparent ensemble; elles chantent et dansent pendant les travaux champêtres; les femmes et les jeunes filles se tressent les unes les autres (occasion de rassemblement.
Les enfants jouent ensemble dans la boue, dans l'eau, et meme dans la saleté. Ils font preuve de créativité sans bornes a partir de la nature; pas besoin de jouets fabriques ou de jeux électroniques pour se divertir.

[[pinterest 44191640063939986]]

[[pinterest 445363850620045339]]

[[youtube d0l1tMp5soc]]

Meme aux occasions de réjouissance et de celebration, dans les danses, on sent la joie de vivre, une joie de vivre contagieuse.
Ah l'Afrique, qu'il y'a la joie de vivre!!

[[amazoncontextwidget]]
`
