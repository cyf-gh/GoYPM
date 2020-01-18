package _ypm_test

import(
	"testing"
	"GoYPM/parse"
)

func Test_SplitHeadBody(t *testing.T) {
	head, body, err := ypm_parse.SplitHeadBody("head@")
	if !( err == nil && body == "" && head == "head" ) {
		t.Fatal( head, body, err )
	}
	head, body, err = ypm_parse.SplitHeadBody("head[@]@")

	if !( err == nil && body == "@" && head == "head" ) {
		t.Fatal( head, body, err )
	}

	head, body, err = ypm_parse.SplitHeadBody("head[@][@]")
	if !( err != nil && body == "" && head == "" ) {
		t.Fatal( head, body, err )
	}

	head, body, err = ypm_parse.SplitHeadBody("head@@")
	if !( err != nil && body == "" && head == "" ) {
		t.Fatal( head, body, err )
	}

	head, body, err = ypm_parse.SplitHeadBody("head[[@]][@]")
	if !( err == nil && body == "[@]" && head == "head" ) {
		t.Fatal( head, body, err )
	}

	head, body, err = ypm_parse.SplitHeadBody("head[[@]][@],[@]")
	if !( err == nil && body == "[@],[@]" && head == "head" ) {
		t.Fatal( head, body, err )
	}
	t.Log("OK")
}

func Test_SplitParaments(t *testing.T) {
	ps := ypm_parse.SplitParaments( "\\\\\\,,param2" );
	if !( ps[0] == "\\," && ps[1] == "param2") {
		t.Fatal( ps )
	}
	ps = ypm_parse.SplitParaments( "\\\\\\,param2" );
	if !( ps[0] == "\\,param2" ) {
		t.Fatal( ps )
	}
	ps = ypm_parse.SplitParaments( "\\\\\\\\,param2" );
	if !( ps[0] == "\\\\" && ps[1] == "param2") {
		t.Fatal( ps )
	}
	t.Log("OK")
}