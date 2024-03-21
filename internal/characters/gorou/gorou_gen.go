// Code generated by "pipeline"; DO NOT EDIT.
package gorou

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
	3: {"travel"},
	7: {"hold", "travel", "weakspot"},
}

func init() {
	base = &model.AvatarData{}
	err := prototext.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
	validation.RegisterCharParamValidationFunc(keys.Gorou, ValidateParamKeys)
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
	attack = [][]float64{
		attack_1,
		attack_2,
		attack_3,
		attack_4,
	}
)

var (
	// attack: aim = [4]
	aim = []float64{
		0.43860000371932983,
		0.47429999709129333,
		0.5099999904632568,
		0.5609999895095825,
		0.5967000126838684,
		0.637499988079071,
		0.6935999989509583,
		0.7497000098228455,
		0.8058000206947327,
		0.8669999837875366,
		0.9282000064849854,
		0.9894000291824341,
		1.0506000518798828,
		1.111799955368042,
		1.1729999780654907,
	}
	// attack: attack_1 = [0]
	attack_1 = []float64{
		0.3775399923324585,
		0.408270001411438,
		0.4390000104904175,
		0.4828999936580658,
		0.5136299729347229,
		0.5487499833106995,
		0.5970399975776672,
		0.645330011844635,
		0.6936200261116028,
		0.7462999820709229,
		0.7989799976348877,
		0.8516600131988525,
		0.9043400287628174,
		0.9570199847221375,
		1.0096999406814575,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.3715200126171112,
		0.4017600119113922,
		0.4320000112056732,
		0.47519999742507935,
		0.5054399967193604,
		0.5400000214576721,
		0.5875200033187866,
		0.6350399851799011,
		0.6825600266456604,
		0.7343999743461609,
		0.7862399816513062,
		0.8380799889564514,
		0.8899199962615967,
		0.9417600035667419,
		0.9936000108718872,
	}
	// attack: attack_3 = [2]
	attack_3 = []float64{
		0.4945000112056732,
		0.5347499847412109,
		0.574999988079071,
		0.6324999928474426,
		0.6727499961853027,
		0.71875,
		0.7820000052452087,
		0.8452500104904175,
		0.9085000157356262,
		0.9775000214576721,
		1.0464999675750732,
		1.1154999732971191,
		1.184499979019165,
		1.253499984741211,
		1.3224999904632568,
	}
	// attack: attack_4 = [3]
	attack_4 = []float64{
		0.589959979057312,
		0.6379799842834473,
		0.6859999895095825,
		0.7545999884605408,
		0.802619993686676,
		0.8575000166893005,
		0.9329599738121033,
		1.0084199905395508,
		1.0838799476623535,
		1.166200041770935,
		1.248520016670227,
		1.330839991569519,
		1.413159966468811,
		1.495479941368103,
		1.5778000354766846,
	}
	// attack: fullaim = [5]
	fullaim = []float64{
		1.2400000095367432,
		1.3329999446868896,
		1.4259999990463257,
		1.5499999523162842,
		1.6430000066757202,
		1.7359999418258667,
		1.8600000143051147,
		1.9839999675750732,
		2.1080000400543213,
		2.2320001125335693,
		2.3559999465942383,
		2.4800000190734863,
		2.634999990463257,
		2.7899999618530273,
		2.944999933242798,
	}
	// skill: skill = [0]
	skill = []float64{
		1.0720000267028809,
		1.152400016784668,
		1.232800006866455,
		1.340000033378601,
		1.4204000234603882,
		1.5008000135421753,
		1.6080000400543213,
		1.7151999473571777,
		1.8223999738693237,
		1.9296000003814697,
		2.036799907684326,
		2.1440000534057617,
		2.2780001163482666,
		2.4119999408721924,
		2.5460000038146973,
	}
	// skill: skillDefBonus = [1]
	skillDefBonus = []float64{
		206.16000366210938,
		221.6219940185547,
		237.08399963378906,
		257.70001220703125,
		273.1619873046875,
		288.6239929199219,
		309.239990234375,
		329.8559875488281,
		350.47198486328125,
		371.0880126953125,
		391.7040100097656,
		412.32000732421875,
		438.0899963378906,
		463.8599853515625,
		489.6300048828125,
	}
	// burst: burst = [0]
	burst = []float64{
		0.9821599721908569,
		1.0558220148086548,
		1.1294840574264526,
		1.2276999950408936,
		1.3013620376586914,
		1.3750239610671997,
		1.4732400178909302,
		1.571455955505371,
		1.6696720123291016,
		1.7678879499435425,
		1.866104006767273,
		1.9643199443817139,
		2.087090015411377,
		2.20986008644104,
		2.332629919052124,
	}
	// burst: burstTick = [1]
	burstTick = []float64{
		0.6129999756813049,
		0.6589750051498413,
		0.7049499750137329,
		0.7662500143051147,
		0.8122249841690063,
		0.8582000136375427,
		0.9194999933242798,
		0.9807999730110168,
		1.042099952697754,
		1.1033999919891357,
		1.1647000312805176,
		1.2259999513626099,
		1.3026249408721924,
		1.3792500495910645,
		1.455875039100647,
	}
)
