package Talk

import "math/rand"
import "time"

var TimeList []string = []string{"前天", "昨天", "今天", "明天", "后天", "刚刚", "有一天", "那天", "忘了是什么时候了"}

var PlaceList []string = []string{
	"公园", "超市", "家", "学校",
	"车站前", "路边", "湖边", "食堂",
	"厕所", "忘记是哪里了", "操场", "图书馆",
}

var ObjectList []string = []string{
	"小猫(/·ω·/)", "果汁", "树叶", "200元钱$_$",
	"梳子", "肥皂", "它", "皮球",
	"鞋子", "糖糖", "钥匙", "什么来着？嘛...",
}

var SubjectList []string = []string{
	"我", "偶", "窝", "萤萤",
	"小萤", "你知道是谁的吧^_^,",
}

var AdjList []string = []string{
	"可爱的", "有点生气的", "调皮的", "神气的", "机智的", "有魅力的",
	"可萌可萌的", "高兴的", "小小的", "神奇的", "形容不清楚的", "兴奋的",
}

var VerbList0 []string = []string{
	"捡到了", "弄丢了", "认识了", "看到了", "发现了", "遇到了",
	"表扬了", "赞美了", "批评了",
}

var FinList []string = []string{
	"真是太机智了", "赞~", "喵=(①ω①)=", "真是觉得可厉害呢☆",
	"哇哇!!", "噗~嘿★", "咪啪~", "嗯，嗯，就是这样^_^",
	"哼~", "呜咕^",
}

/*


*/
var VerbList1 []string = []string{
	"捡起来", "扔出去", "抛弃", "嘲笑",
	"鄙视", "摸头", "表扬", "批评",
	"夸奖", "讨厌", "赞", "发现",
}

/*



*/
func F1() string {

	s := ""
	a := int64(time.Now().Nanosecond())
	rand.Seed(a)
	s += "就是"
	s += TimeList[rand.Intn(len(TimeList))]
	s += ", "
	s += SubjectList[rand.Intn(len(SubjectList))]
	s += "在"
	rand.Seed(int64(time.Now().Nanosecond()))
	s += PlaceList[rand.Intn(len(PlaceList))]
	s += VerbList0[rand.Intn(len(VerbList0))]
	s += AdjList[rand.Intn(len(AdjList))]
	rand.Seed(int64(time.Now().Nanosecond()))
	s += ObjectList[rand.Intn(len(ObjectList))]
	s += "~  "
	s += FinList[rand.Intn(len(FinList))]
	return s
}

//
func F2() string {
	s := ""
	a := int64(time.Now().Nanosecond())
	rand.Seed(a)
	s += "大概是"
	s += TimeList[rand.Intn(len(TimeList))]
	s += "吧, "
	s += SubjectList[rand.Intn(len(SubjectList))]
	s += "在"
	rand.Seed(int64(time.Now().Nanosecond()))
	s += PlaceList[rand.Intn(len(PlaceList))]
	s += "被"
	s += AdjList[rand.Intn(len(AdjList))]
	rand.Seed(int64(time.Now().Nanosecond()))
	s += ObjectList[rand.Intn(len(ObjectList))]
	s += VerbList1[rand.Intn(len(VerbList1))]
	s += "了"
	s += ">_<  "
	s += FinList[rand.Intn(len(FinList))]
	return s
}
func F3() string {
	s := ""
	a := int64(time.Now().Nanosecond())
	rand.Seed(a)
	s += "呐, 呐, "
	s += TimeList[rand.Intn(len(TimeList))]
	s += "吧, "
	s += ObjectList[rand.Intn(len(ObjectList))]
	s += "在"
	rand.Seed(int64(time.Now().Nanosecond()))
	s += PlaceList[rand.Intn(len(PlaceList))]
	s += "被"
	s += AdjList[rand.Intn(len(AdjList))]
	rand.Seed(int64(time.Now().Nanosecond()))
	s += SubjectList[rand.Intn(len(SubjectList))]
	s += VerbList1[rand.Intn(len(VerbList1))]
	s += "了"
	s += ">_<  "
	s += FinList[rand.Intn(len(FinList))]
	return s
}

func F() string {
	rand.Seed(int64(time.Now().Nanosecond()))
	a := rand.Intn(2)
	if a == 0 {
		return F1()
	} else {
		a = rand.Intn(2)
		if a == 0 {
			return F2()
		} else {
			return F3()
		}
	}
}
