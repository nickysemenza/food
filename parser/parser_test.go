package parser

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func Test_parse(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	res, err := Parse(context.Background(), "1.2 c (37gr) all p-purpose flour")
	require.NoError(t, err)
	spew.Dump(res)
	res, err = Parse(context.Background(), "1 ½ teaspoons/6 grams baking powder")
	require.NoError(t, err)
	spew.Dump(res)
	res, err = Parse(context.Background(), "3 tablespoons/43.2 grams unsalted butter, melted")
	require.NoError(t, err)
	spew.Dump(res)
	t.Fail()
}

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    *Ingredient
		wantErr bool
	}{
		{name: "empty", wantErr: true},
		{name: "malformed 1", s: "flour", wantErr: true},
		{name: "malformed 2", s: "cup", wantErr: true},
		{name: "malformed 3", s: "1", wantErr: true},
		{name: "malformed 3", s: " 1 1/e2 cup", wantErr: true},
		{name: "malformed 3", s: "\a", wantErr: true},
		{name: "basic vol", s: "1 cup flour", want: &Ingredient{Name: "flour", Volume: Measurement{Value: 1, Unit: "cup"}}},
		{name: "basic vol decimal", s: "1.2 cup flour", want: &Ingredient{Name: "flour", Volume: Measurement{Value: 1.2, Unit: "cup"}}},
		{name: "basic vol fraction", s: "1 ½ cup flour", want: &Ingredient{Name: "flour", Volume: Measurement{Value: 1.5, Unit: "cup"}}},
		{name: "basic mass", s: "100 grams flour", want: &Ingredient{Name: "flour", Weight: Measurement{Value: 100, Unit: "grams"}}},
		{name: "basic mass", s: "100 gr flour", want: &Ingredient{Name: "flour", Weight: Measurement{Value: 100, Unit: "gr"}}},
		{name: "basic mass and vol", s: "1/2 cup (60 grams) flour", want: &Ingredient{Name: "flour",
			Weight: Measurement{Value: 60, Unit: "grams"},
			Volume: Measurement{Value: 0.5, Unit: "cup"},
		}},
		{name: "mass and vol, modifier", s: "¾ cup (90 grams) flour, sifted", want: &Ingredient{Name: "flour",
			Weight:   Measurement{Value: 90, Unit: "grams"},
			Volume:   Measurement{Value: 0.75, Unit: "cup"},
			Modifier: "sifted",
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(context.Background(), tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				// return
			}
			require.Equal(t, tt.want, got)
		})
	}
}
