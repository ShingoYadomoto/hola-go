package hola_go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHupaiCalculater_peko(t *testing.T) {
	type fields struct {
		standard StandardHoluPattern
	}
	tests := []struct {
		name   string
		fields fields
		want   []HandType
	}{
		{
			name: "true case. ipeko case",
			fields: fields{
				standard: StandardHoluPattern{
					Mentsu: []mentsu{
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{中, 中, 中}},
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{白, 白, 白}},
					},
				},
			},
			want: []HandType{一盃口},
		},
		{
			name: "false case. ipeko case. 3 same shuntsu",
			fields: fields{
				standard: StandardHoluPattern{
					Mentsu: []mentsu{
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{白, 白, 白}},
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{一萬, 二萬, 三萬}},
					},
				},
			},
			want: []HandType{一盃口},
		},
		{
			name: "true case. ryanpeko case",
			fields: fields{
				standard: StandardHoluPattern{
					Mentsu: []mentsu{
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{七萬, 八萬, 九萬}},
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{七萬, 八萬, 九萬}},
					},
				},
			},
			want: []HandType{二盃口},
		},
		{
			name: "false case. ipeko case. 3 same shuntsu",
			fields: fields{
				standard: StandardHoluPattern{
					Mentsu: []mentsu{
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{東, 東, 東}},
						{pais: []pai{中, 中, 中}},
						{pais: []pai{白, 白, 白}},
					},
				},
			},
			want: []HandType{},
		},
		{
			name: "false case. fulou case",
			fields: fields{
				standard: StandardHoluPattern{
					Mentsu: []mentsu{
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{中, 中, 中}},
						{pais: []pai{一萬, 二萬, 三萬}},
					},
					FulouMentsu: FulouMentsuList{
						{mentsu: mentsu{pais: []pai{白, 白, 白}}},
					},
				},
			},
			want: []HandType{},
		},
		{
			name: "false case. not same shuntsu case",
			fields: fields{
				standard: StandardHoluPattern{
					Mentsu: []mentsu{
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{東, 東, 東}},
						{pais: []pai{中, 中, 中}},
						{pais: []pai{白, 白, 白}},
					},
				},
			},
			want: []HandType{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hc := HupaiCalculater{
				standard: tt.fields.standard,
			}
			assert.Equalf(t, tt.want, hc.peko(), "ipeko()")
		})
	}
}

func TestFullHupaiCalculater_Hupai(t *testing.T) {
	type fields struct {
		fullParrern FullHoluPattern
		zhuangfeng  Zhuangfeng
		zifeng      Zifeng
		isTsumo     bool
	}
	tests := []struct {
		name   string
		fields fields
		want   PossibleAllHands
	}{
		{
			name: "役なし",
			fields: fields{
				fullParrern: FullHoluPattern{
					Standard: []StandardHoluPattern{
						{
							Mentsu: []mentsu{
								{pais: []pai{四萬, 五萬, 六萬}, holuPai: &四萬},
								{pais: []pai{三筒, 四筒, 五筒}},
								{pais: []pai{五索, 六索, 七索}},
								{pais: []pai{九索, 九索, 九索}},
							},
							Head:    二索,
							HoluPai: 四萬,
						},
					},
					IsTsumo: false,
				},
				zhuangfeng: 東場,
				zifeng:     東家,
				isTsumo:    false,
			},
			want: []AllHands{{}},
		},

		{
			name: "メンタンピン",
			fields: fields{
				fullParrern: FullHoluPattern{
					Standard: []StandardHoluPattern{
						{
							Mentsu: []mentsu{
								{pais: []pai{四萬, 五萬, 六萬}, holuPai: &四萬},
								{pais: []pai{三筒, 四筒, 五筒}},
								{pais: []pai{五索, 六索, 七索}},
								{pais: []pai{二索, 三索, 四索}},
							},
							Head:    二索,
							HoluPai: 四萬,
						},
					},
					IsTsumo: true,
				},
				zhuangfeng: 東場,
				zifeng:     東家,
				isTsumo:    true,
			},
			want: []AllHands{{門前清自摸和, 平和, 断幺九}},
		},

		{
			name: "翻牌場風",
			fields: fields{
				fullParrern: FullHoluPattern{
					Standard: []StandardHoluPattern{
						{
							Mentsu: []mentsu{
								{pais: []pai{七索, 八索, 九索}, holuPai: &七索},
								{pais: []pai{三筒, 四筒, 五筒}},
								{pais: []pai{七筒, 八筒, 九筒}},
								{pais: []pai{東, 東, 東}},
							},
							Head:    四索,
							HoluPai: 七索,
						},
					},
				},
				zhuangfeng: 東場,
				zifeng:     南家,
			},
			want: []AllHands{{翻牌場風}},
		},

		{
			name: "翻牌自風",
			fields: fields{
				fullParrern: FullHoluPattern{
					Standard: []StandardHoluPattern{
						{
							Mentsu: []mentsu{
								{pais: []pai{七索, 八索, 九索}, holuPai: &七索},
								{pais: []pai{三筒, 四筒, 五筒}},
								{pais: []pai{七筒, 八筒, 九筒}},
								{pais: []pai{東, 東, 東}},
							},
							Head:    四索,
							HoluPai: 七索,
						},
					},
				},
				zhuangfeng: 南場,
				zifeng:     東家,
			},
			want: []AllHands{{翻牌自風}},
		},

		{
			name: "翻牌白",
			fields: fields{
				fullParrern: FullHoluPattern{
					Standard: []StandardHoluPattern{
						{
							Mentsu: []mentsu{
								{pais: []pai{七索, 八索, 九索}, holuPai: &七索},
								{pais: []pai{三筒, 四筒, 五筒}},
								{pais: []pai{七筒, 八筒, 九筒}},
								{pais: []pai{白, 白, 白}},
							},
							Head:    四索,
							HoluPai: 七索,
						},
					},
				},
				zhuangfeng: 南場,
				zifeng:     東家,
			},
			want: []AllHands{{翻牌白}},
		},

		{
			name: "翻牌發",
			fields: fields{
				fullParrern: FullHoluPattern{
					Standard: []StandardHoluPattern{
						{
							Mentsu: []mentsu{
								{pais: []pai{七索, 八索, 九索}, holuPai: &七索},
								{pais: []pai{三筒, 四筒, 五筒}},
								{pais: []pai{七筒, 八筒, 九筒}},
								{pais: []pai{發, 發, 發}},
							},
							Head:    四索,
							HoluPai: 七索,
						},
					},
				},
				zhuangfeng: 南場,
				zifeng:     東家,
			},
			want: []AllHands{{翻牌發}},
		},

		{
			name: "翻牌中",
			fields: fields{
				fullParrern: FullHoluPattern{
					Standard: []StandardHoluPattern{
						{
							Mentsu: []mentsu{
								{pais: []pai{七索, 八索, 九索}, holuPai: &七索},
								{pais: []pai{三筒, 四筒, 五筒}},
								{pais: []pai{七筒, 八筒, 九筒}},
								{pais: []pai{中, 中, 中}},
							},
							Head:    四索,
							HoluPai: 七索,
						},
					},
				},
				zhuangfeng: 南場,
				zifeng:     東家,
			},
			want: []AllHands{{翻牌中}},
		},

		{
			name: "一盃口",
			fields: fields{
				fullParrern: FullHoluPattern{
					Standard: []StandardHoluPattern{
						{
							Mentsu: []mentsu{
								{pais: []pai{一萬, 二萬, 三萬}},
								{pais: []pai{四索, 四索, 四索}},
								{pais: []pai{一萬, 二萬, 三萬}},
								{pais: []pai{三索, 三索, 三索}},
							},
							Head:    中,
							HoluPai: 中,
						},
					},
				},
				zhuangfeng: 南場,
				zifeng:     東家,
			},
			want: []AllHands{{一盃口}},
		},

		{
			name: "七対子・二盃口",
			fields: fields{
				fullParrern: FullHoluPattern{
					Standard: []StandardHoluPattern{
						{
							Mentsu: []mentsu{
								{pais: []pai{一萬, 二萬, 三萬}},
								{pais: []pai{一萬, 二萬, 三萬}},
								{pais: []pai{七索, 八索, 九索}},
								{pais: []pai{七索, 八索, 九索}},
							},
							Head:    一萬,
							HoluPai: 一萬,
						},
					},
					Titoitsu: &TitoitsuHoluPattern{
						Menzen: []pai{
							一萬, 二萬, 三萬,
							七索, 八索, 九索,
						},
						HoluPai: 一萬,
					},
				},
				zhuangfeng: 南場,
				zifeng:     東家,
			},
			want: []AllHands{
				{七対子},
				{二盃口},
			},
		},

		{name: "三色同順"},
		{name: "一気通貫"},
		{name: "混全帯幺九"},
		{name: "七対子"},
		{name: "対々和"},
		{name: "三暗刻"},
		{name: "三槓子"},
		{name: "三色同刻"},
		{name: "混老頭"},
		{name: "小三元"},
		{name: "混一色"},
		{name: "純全帯幺九"},
		{name: "清一色"},
		{name: "国士無双"},
		{name: "国士無双十三面"},
		{name: "四暗刻"},
		{name: "四暗刻単騎"},
		{name: "大三元"},
		{name: "小四喜"},
		{name: "大四喜"},
		{name: "字一色"},
		{name: "緑一色"},
		{name: "清老頭"},
		{name: "四槓子"},
		{name: "九蓮宝燈"},
		{name: "純正九蓮宝燈"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fhc := FullHupaiCalculater{
				fullParrern: tt.fields.fullParrern,
				zhuangfeng:  tt.fields.zhuangfeng,
				zifeng:      tt.fields.zifeng,
				isTsumo:     tt.fields.isTsumo,
			}
			assert.Equalf(t, tt.want.String(), fhc.Hupai().String(), "Hupai()")
		})
	}
}
