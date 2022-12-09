package hand

type HandType int

const (
	門前清自摸和 HandType = iota + 1
	場風
	自風
	白
	發
	中
	平和
	断幺九
	一盃口
	三色同順
	一気通貫
	混全帯幺九
	七対子
	対々和
	三暗刻
	三槓子
	三色同刻
	混老頭
	小三元
	混一色
	純全帯幺九
	二盃口
	清一色
)

type Fanshu int

func (ht HandType) Fanshu(isMenzen bool) Fanshu {
	var fanshu Fanshu

	switch ht {
	case 門前清自摸和, 場風, 自風, 白, 發, 中, 平和, 断幺九, 一盃口:
		fanshu = 1
	case 三色同順, 一気通貫, 混全帯幺九, 七対子, 対々和, 三暗刻, 三槓子, 三色同刻, 混老頭, 小三元, 混一色:
		fanshu = 2
	case 純全帯幺九, 二盃口:
		fanshu = 3
	case 清一色:
		fanshu = 6
	}

	if isMenzen {
		return fanshu
	}

	switch ht {
	case 三色同順, 一気通貫, 混全帯幺九:
		fanshu = 1
	case 混一色, 純全帯幺九:
		fanshu = 2
	case 清一色:
		fanshu = 5
	}

	return fanshu
}

type zhuangfeng int

const (
	東家 zhuangfeng = iota + 1
	西家
	南家
	北家
)

type zifeng int

const (
	東場 zifeng = iota + 1
	西場
	南場
	北場
)
