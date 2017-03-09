package formatter

import (
	"github.com/mtso/exercises/util"
	"testing"
)

/*
In:
        // The spectacle before us was indeed sublime.
        // Apparently we had reached a great height in the atmosphere, for the sky was a dead black, and the stars had ceased to twinkle. By the same illusion which lifts the horizon of the sea to the level of the spectator on a hillside, the sable cloud beneath was dished out, and the car seemed to float in the middle of an immense dark sphere, whose upper half was strewn with silver.
Out:
        // The spectacle before us was indeed sublime.
        // Apparently we had reached a great height in the atmosphere, for the
        // sky was a dead black, and the stars had ceased to twinkle. By the
        // same illusion which lifts the horizon of the sea to the level of the
        // spectator on a hillside, the sable cloud beneath was dished out, and
        // the car seemed to float in the middle of an immense dark sphere,
        // whose upper half was strewn with silver.
*/

var sampleComments = [...]struct {
	in  []byte
	out []byte
}{
	{
		[]byte("        // The spectacle before us was indeed sublime.\n        // Apparently we had reached a great height in the atmosphere, for the sky was a dead black, and the stars had ceased to twinkle. By the same illusion which lifts the horizon of the sea to the level of the spectator on a hillside, the sable cloud beneath was dished out, and the car seemed to float in the middle of an immense dark sphere, whose upper half was strewn with silver."),
		[]byte("        // The spectacle before us was indeed sublime.\n        // Apparently we had reached a great height in the atmosphere, for the\n        // sky was a dead black, and the stars had ceased to twinkle. By the\n        // same illusion which lifts the horizon of the sea to the level of the\n        // spectator on a hillside, the sable cloud beneath was dished out, and\n        // the car seemed to float in the middle of an immense dark sphere,\n        // whose upper half was strewn with silver."),
	},
}

func TestFormatter(t *testing.T) {
	formatter := NewFormatter(80)
	actual := formatter.fmt(sampleComments[0].in)

	if !util.AssertEqualBytes(actual, sampleComments[0].out) {
		t.Errorf("%q != %q", expected, sampleComments[0])
	}
}
