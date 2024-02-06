// Code generated by "pipeline"; DO NOT EDIT.
package sara

import (
	_ "embed"

	"github.com/genshinsim/gcsim/pkg/model"
	"google.golang.org/protobuf/encoding/prototext"
)

//go:embed data_gen.textproto
var pbData []byte
var base *model.AvatarData

func init() {
	base = &model.AvatarData{}
	err := prototext.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
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
		attack_5,
	}
)

var (
	// attack: aim = [5]
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
		0.36893999576568604,
		0.39897000789642334,
		0.42899999022483826,
		0.47189998626708984,
		0.5019299983978271,
		0.5362499952316284,
		0.5834400057792664,
		0.6306300163269043,
		0.6778200268745422,
		0.7293000221252441,
		0.780780017375946,
		0.832260012626648,
		0.8837400078773499,
		0.9352200031280518,
		0.9866999983787537,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.3869999945163727,
		0.41850000619888306,
		0.44999998807907104,
		0.4950000047683716,
		0.5264999866485596,
		0.5625,
		0.6119999885559082,
		0.6614999771118164,
		0.7110000252723694,
		0.7649999856948853,
		0.8190000057220459,
		0.8730000257492065,
		0.9269999861717224,
		0.9810000061988831,
		1.034999966621399,
	}
	// attack: attack_3 = [2]
	attack_3 = []float64{
		0.48504000902175903,
		0.5245199799537659,
		0.5640000104904175,
		0.6204000115394592,
		0.6598799824714661,
		0.7049999833106995,
		0.7670400142669678,
		0.8290799856185913,
		0.8911200165748596,
		0.9588000178337097,
		1.026479959487915,
		1.0941599607467651,
		1.1618399620056152,
		1.2295199632644653,
		1.2971999645233154,
	}
	// attack: attack_4 = [3]
	attack_4 = []float64{
		0.5039600133895874,
		0.544979989528656,
		0.5860000252723694,
		0.644599974155426,
		0.6856200098991394,
		0.7325000166893005,
		0.7969599962234497,
		0.8614199757575989,
		0.9258800148963928,
		0.9962000250816345,
		1.0665199756622314,
		1.1368399858474731,
		1.2071599960327148,
		1.2774800062179565,
		1.3478000164031982,
	}
	// attack: attack_5 = [4]
	attack_5 = []float64{
		0.5805000066757202,
		0.6277499794960022,
		0.675000011920929,
		0.7425000071525574,
		0.7897499799728394,
		0.84375,
		0.9179999828338623,
		0.9922500252723694,
		1.066499948501587,
		1.1475000381469727,
		1.2285000085830688,
		1.309499979019165,
		1.3904999494552612,
		1.471500039100647,
		1.5525000095367432,
	}
	// attack: fullaim = [6]
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
	// skill: atkBuff = [1]
	atkBuff = []float64{
		0.4296000003814697,
		0.46182000637054443,
		0.49404001235961914,
		0.5370000004768372,
		0.5692200064659119,
		0.6014400124549866,
		0.6444000005722046,
		0.6873599886894226,
		0.7303199768066406,
		0.7732800245285034,
		0.8162400126457214,
		0.8592000007629395,
		0.9128999710083008,
		0.9666000008583069,
		1.020300030708313,
	}
	// skill: skill = [0]
	skill = []float64{
		1.257599949836731,
		1.3519200086593628,
		1.446239948272705,
		1.5720000267028809,
		1.6663199663162231,
		1.760640025138855,
		1.8863999843597412,
		2.012160062789917,
		2.1379199028015137,
		2.2636799812316895,
		2.3894400596618652,
		2.515199899673462,
		2.6723999977111816,
		2.8296000957489014,
		2.986799955368042,
	}
	// burst: burstCluster = [1]
	burstCluster = []float64{
		0.34119999408721924,
		0.3667899966239929,
		0.3923799991607666,
		0.42649999260902405,
		0.45208999514579773,
		0.4776799976825714,
		0.5117999911308289,
		0.5459200143814087,
		0.5800399780273438,
		0.6141600012779236,
		0.6482800245285034,
		0.6823999881744385,
		0.7250499725341797,
		0.7677000164985657,
		0.8103500008583069,
	}
	// burst: burstMain = [0]
	burstMain = []float64{
		4.0960001945495605,
		4.403200149536133,
		4.710400104522705,
		5.119999885559082,
		5.427199840545654,
		5.734399795532227,
		6.144000053405762,
		6.553599834442139,
		6.963200092315674,
		7.372799873352051,
		7.782400131225586,
		8.192000389099121,
		8.704000473022461,
		9.215999603271484,
		9.727999687194824,
	}
)
