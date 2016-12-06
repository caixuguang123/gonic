package models
import "errors"

type Article struct {
	ID      int `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = [] Article{
	{ID:1, Title:`再别康桥


	`, Content:`
轻轻的我走了，
正如我轻轻的来；
我轻轻的招手，
作别西天的云彩。
----
那河畔的金柳，
是夕阳中的新娘；
波光里的艳影，
在我的心头荡漾。
----
软泥上的青荇⑴，
油油的在水底招摇⑵；
在康河的柔波里，
我甘心做一条水草！
----
那榆荫下的一潭，
不是清泉，是天上虹；
揉碎在浮藻间，
沉淀着彩虹似的梦。
----
寻梦？撑一支长篙⑶，
向青草更青处漫溯⑷；
满载一船星辉，
在星辉斑斓里放歌。
----
但我不能放歌，
悄悄是别离的笙箫；
夏虫也为我沉默，
沉默是今晚的康桥！
----
悄悄的我走了，
正如我悄悄的来；
我挥一挥衣袖，
不带走一片云彩。`},
	{ID:2, Title:`致橡树

	`, Content:`
   我如果爱你——
　　绝不像攀援①的凌霄花，
　　借你的高枝炫耀自己；
　　我如果爱你——
　　绝不学痴情的鸟儿，
　　为绿荫重复单调的歌曲；
　　也不止像泉源，
　　常年送来清凉的慰藉②；
　　也不止像险峰，
　　增加你的高度，衬托你的威仪。
　　甚至日光，
　　甚至春雨。
   不，这些都还不够！
　　我必须是你近旁的一株木棉，
　　作为树的形象和你站在一起。
　　根，紧握在地下；
　　叶，相触在云里。
　　每一阵风过，
　　我们都互相致意，
　　但没有人，
　　听懂我们的言语。
　　你有你的铜枝铁干，
　　像刀，像剑，也像戟；
　　我有我红硕的花朵，
　　像沉重的叹息，
　　又像英勇的火炬。
    我们分担寒潮、风雷、霹雳；
　　我们共享雾霭③、流岚⑤、虹霓④。
　　仿佛永远分离，
　　却又终身相依。
　　这才是伟大的爱情，
　　坚贞就在这里：
　　爱——
　　不仅爱你伟岸的身躯，
　　也爱你坚持的位置，
　　足下的土地。`},
	{
		ID:3,
		Title:"长相思 其一",
		Content:`长相思，在长安。
络纬秋啼金井阑，微霜凄凄簟色寒。
孤灯不明思欲绝，卷帷望月空长叹。
美人如花隔云端！
上有青冥之长天，下有渌水之波澜。
天长路远魂飞苦，梦魂不到关山难。
长相思，摧心肝！`,
	},
}

func GetAllArticles() []Article {
	return articleList
}

func GetArticleById(id int) (*Article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

func CreateNewArticle(title, content string) (*Article, error) {
	id := len(articleList) + 1;
	articleNew := Article{id, title, content}
	articleList = append(articleList, articleNew)
	return &articleNew, nil
}
