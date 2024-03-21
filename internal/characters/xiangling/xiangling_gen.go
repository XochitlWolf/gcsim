// Code generated by "pipeline"; DO NOT EDIT.
package xiangling

import (
	_ "embed"

	"fmt"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/keys"
	"github.com/genshinsim/gcsim/pkg/gcs/validation"
	"github.com/genshinsim/gcsim/pkg/model"
	"google.golang.org/protobuf/encoding/prototext"
	"slices"
)

//go:embed data_gen.textproto
var pbData []byte
var base *model.AvatarData
var paramKeysValidation = map[action.Action][]string{
	1: {"a4_delay"},
	5: {"collision"},
	6: {"collision"},
}

func init() {
	base = &model.AvatarData{}
	err := prototext.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
	validation.RegisterCharParamValidationFunc(keys.Xiangling, ValidateParamKeys)
}

func ValidateParamKeys(a action.Action, keys []string) error {
	valid, ok := paramKeysValidation[a]
	if !ok {
		return nil
	}
	for _, v := range keys {
		if !slices.Contains(valid, v) {
			return fmt.Errorf("key %v is invalid for action %v", v, a.String())
		}
	}
	return nil
}

func (x *char) Data() *model.AvatarData {
	return base
}

var (
	attack = [][][]float64{
		{attack_1},
		{attack_2},
		attack_3,
		attack_4,
		{attack_5},
	}
)

var (
	// attack: attack_1 = [0]
	attack_1 = []float64{
		0.4205400049686432,
		0.4547699987888336,
		0.48899999260902405,
		0.5378999710083008,
		0.572130024433136,
		0.6112499833106995,
		0.6650400161743164,
		0.7188299894332886,
		0.7726200222969055,
		0.8313000202178955,
		0.8985369801521301,
		0.9776089787483215,
		1.0566799640655518,
		1.1357510089874268,
		1.2220109701156616,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.4214000105857849,
		0.45570001006126404,
		0.49000000953674316,
		0.5389999747276306,
		0.5733000040054321,
		0.612500011920929,
		0.6664000153541565,
		0.720300018787384,
		0.7742000222206116,
		0.8330000042915344,
		0.9003750085830688,
		0.9796079993247986,
		1.0588409900665283,
		1.1380740404129028,
		1.2245099544525146,
	}
	// attack: attack_3 = [2 2]
	attack_3 = [][]float64{
		{
			0.26058000326156616,
			0.28178998827934265,
			0.30300000309944153,
			0.33329999446868896,
			0.35451000928878784,
			0.3787499964237213,
			0.41207998991012573,
			0.44541001319885254,
			0.47874000668525696,
			0.5151000022888184,
			0.5567619800567627,
			0.605758011341095,
			0.6547530293464661,
			0.7037479877471924,
			0.7571970224380493,
		},
		{
			0.26058000326156616,
			0.28178998827934265,
			0.30300000309944153,
			0.33329999446868896,
			0.35451000928878784,
			0.3787499964237213,
			0.41207998991012573,
			0.44541001319885254,
			0.47874000668525696,
			0.5151000022888184,
			0.5567619800567627,
			0.605758011341095,
			0.6547530293464661,
			0.7037479877471924,
			0.7571970224380493,
		},
	}
	// attack: attack_4 = [3 3 3 3]
	attack_4 = [][]float64{
		{
			0.14103999733924866,
			0.15252000093460083,
			0.164000004529953,
			0.18039999902248383,
			0.191880002617836,
			0.20499999821186066,
			0.22303999960422516,
			0.24108000099658966,
			0.25911998748779297,
			0.27880001068115234,
			0.3013499975204468,
			0.3278689980506897,
			0.3543879985809326,
			0.38090598583221436,
			0.4098359942436218,
		},
		{
			0.14103999733924866,
			0.15252000093460083,
			0.164000004529953,
			0.18039999902248383,
			0.191880002617836,
			0.20499999821186066,
			0.22303999960422516,
			0.24108000099658966,
			0.25911998748779297,
			0.27880001068115234,
			0.3013499975204468,
			0.3278689980506897,
			0.3543879985809326,
			0.38090598583221436,
			0.4098359942436218,
		},
		{
			0.14103999733924866,
			0.15252000093460083,
			0.164000004529953,
			0.18039999902248383,
			0.191880002617836,
			0.20499999821186066,
			0.22303999960422516,
			0.24108000099658966,
			0.25911998748779297,
			0.27880001068115234,
			0.3013499975204468,
			0.3278689980506897,
			0.3543879985809326,
			0.38090598583221436,
			0.4098359942436218,
		},
		{
			0.14103999733924866,
			0.15252000093460083,
			0.164000004529953,
			0.18039999902248383,
			0.191880002617836,
			0.20499999821186066,
			0.22303999960422516,
			0.24108000099658966,
			0.25911998748779297,
			0.27880001068115234,
			0.3013499975204468,
			0.3278689980506897,
			0.3543879985809326,
			0.38090598583221436,
			0.4098359942436218,
		},
	}
	// attack: attack_5 = [4]
	attack_5 = []float64{
		0.7103599905967712,
		0.7681800127029419,
		0.8259999752044678,
		0.9085999727249146,
		0.9664199948310852,
		1.0325000286102295,
		1.12336003780365,
		1.2142200469970703,
		1.3050800561904907,
		1.4041999578475952,
		1.517775058746338,
		1.651339054107666,
		1.7849030494689941,
		1.9184679985046387,
		2.064173936843872,
	}
	// attack: collision = [7]
	collision = []float64{
		0.6393240094184875,
		0.6913620233535767,
		0.743399977684021,
		0.8177400231361389,
		0.8697779774665833,
		0.9292500019073486,
		1.011023998260498,
		1.0927979946136475,
		1.1745719909667969,
		1.2637799978256226,
		1.3529880046844482,
		1.442196011543274,
		1.5314040184020996,
		1.6206120252609253,
		1.709820032119751,
	}
	// attack: highPlunge = [9]
	highPlunge = []float64{
		1.59676194190979,
		1.7267309427261353,
		1.8566999435424805,
		2.042370080947876,
		2.1723389625549316,
		2.3208749294281006,
		2.5251119136810303,
		2.72934889793396,
		2.9335858821868896,
		3.1563899517059326,
		3.3791940212249756,
		3.6019980907440186,
		3.8248019218444824,
		4.047605991363525,
		4.270410060882568,
	}
	// attack: lowPlunge = [8]
	lowPlunge = []float64{
		1.2783770561218262,
		1.3824310302734375,
		1.4864850044250488,
		1.635133981704712,
		1.7391870021820068,
		1.858106017112732,
		2.021620035171509,
		2.1851329803466797,
		2.3486459255218506,
		2.527024984359741,
		2.7054030895233154,
		2.8837809562683105,
		3.0621590614318848,
		3.24053692817688,
		3.418915033340454,
	}
	// attack: nc = [5]
	nc = []float64{
		1.2168999910354614,
		1.3159500360488892,
		1.4149999618530273,
		1.55649995803833,
		1.6555500030517578,
		1.7687499523162842,
		1.924399971961975,
		2.080049991607666,
		2.2356998920440674,
		2.4054999351501465,
		2.6000618934631348,
		2.8288679122924805,
		3.0576729774475098,
		3.2864789962768555,
		3.5360848903656006,
	}
	// skill: guobaTick = [0]
	guobaTick = []float64{
		1.1128000020980835,
		1.1962599754333496,
		1.2797199487686157,
		1.3910000324249268,
		1.4744600057601929,
		1.557919979095459,
		1.6691999435424805,
		1.7804800271987915,
		1.891759991645813,
		2.003040075302124,
		2.1143200397491455,
		2.225600004196167,
		2.3647000789642334,
		2.5037999153137207,
		2.642899990081787,
	}
	// burst: pyronadoInitial = [0 1 2]
	pyronadoInitial = [][]float64{
		{
			0.7200000286102295,
			0.7739999890327454,
			0.828000009059906,
			0.8999999761581421,
			0.9539999961853027,
			1.0080000162124634,
			1.0800000429153442,
			1.1519999504089355,
			1.2239999771118164,
			1.2960000038146973,
			1.3680000305175781,
			1.440000057220459,
			1.5299999713897705,
			1.6200000047683716,
			1.7100000381469727,
		},
		{
			0.8799999952316284,
			0.9459999799728394,
			1.0119999647140503,
			1.100000023841858,
			1.1660000085830688,
			1.2319999933242798,
			1.3200000524520874,
			1.4079999923706055,
			1.496000051498413,
			1.5839999914169312,
			1.6720000505447388,
			1.7599999904632568,
			1.8700000047683716,
			1.9800000190734863,
			2.0899999141693115,
		},
		{
			1.0959999561309814,
			1.1782000064849854,
			1.2604000568389893,
			1.3700000047683716,
			1.4522000551223755,
			1.5343999862670898,
			1.6440000534057617,
			1.753600001335144,
			1.8631999492645264,
			1.9728000164031982,
			2.08240008354187,
			2.191999912261963,
			2.3289999961853027,
			2.4660000801086426,
			2.6029999256134033,
		},
	}
	// burst: pyronadoSpin = [3]
	pyronadoSpin = []float64{
		1.1200000047683716,
		1.2039999961853027,
		1.2879999876022339,
		1.399999976158142,
		1.4839999675750732,
		1.5679999589920044,
		1.6799999475479126,
		1.7920000553131104,
		1.9040000438690186,
		2.0160000324249268,
		2.128000020980835,
		2.240000009536743,
		2.380000114440918,
		2.5199999809265137,
		2.6600000858306885,
	}
)
