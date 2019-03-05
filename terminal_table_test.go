// @author      Liu Yongshuai<liuyongshuai@hotmail.com>
// @date        2018-11-27 20:14

package goUtils

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewTerminalTable(t *testing.T) {
	testStart()
	table := NewTerminalTable()
	table.SetHeader([]string{"index", "shopId", "shopName"})
	table.SetBorderFontColor(ColorType_Blue)
	table.SetRowFontColor(ColorType_Red)
	for _, sinfo := range tmpForTestShopList {
		tmp := strings.Split(sinfo, "|")
		table.AddRow(tmp)
	}
	fmt.Println(table.Render())
	testEnd()
}

func TestNewTerminalTable1(t *testing.T) {
	testStart()
	table := NewTerminalTable()
	table.SetHeader([]string{"index", "shopId", "shopName"})
	for _, sinfo := range tmpForTestShopList {
		tmp := strings.Split(sinfo, "|")
		table.AddRow(tmp)
	}
	fmt.Println(table.Render())
	testEnd()
}

var tmpForTestShopList = []string{
	"0|1152921533330227217|地锅一家人（家常炒菜）ａｂｃａ@￥@#%#ｓｄ🎈🎉ｆ我E２３４３４５んエォサ６３＃＄％＾＄＆％＾（＆我“地锅一家人（家常炒菜）ａｂｃａ@￥@#%#ｓｄ🎈🎉ｆ我E２３４３４５んエォサ６３＃＄％＾＄＆％＾（＆我“地锅一家人（家常炒菜）ａｂｃａ@￥@#%#ｓｄ🎈🎉ｆ我E２３４３４５んエォサ６３＃＄％＾＄＆％＾（＆我“地锅一家人（家常炒菜）ａｂｃａ@￥@#%#ｓｄ🎈🎉ｆ我E２３４３４５んエォサ６３＃＄％＾＄＆％＾（＆我“地锅一家人（家常炒菜）ａｂｃａ@￥@#%#ｓｄ🎈🎉ｆ我E２３４３４５んエォサ６３＃＄％＾＄＆％＾（＆我“地锅一家人（家常炒菜）ａｂｃａ@￥@#%#ｓｄ🎈🎉ｆ我E２３４３４５んエォサ６３＃＄％＾＄＆％＾（＆我“地锅一家人（家常炒菜）ａｂｃａ@￥@#%#ｓｄ🎈🎉ｆ我E２３４３４５んエォサ６３＃＄％＾＄＆％＾（＆我“地锅一家人（家常炒菜）ａｂｃａ@￥@#%#ｓｄ🎈🎉ｆ我E２３４３４５んエォサ６３＃＄％＾＄＆％＾（＆我“地锅一家人（家常炒菜）ａｂｃａ@￥@#%#ｓｄ🎈🎉ｆ我E２３４３４５んエォサ６３＃＄％＾＄＆％＾（＆我“地锅一家人（家常炒菜）ａｂｃａ@￥@#%#ｓｄ🎈🎉ｆ我E２３４３４５んエォサ６３＃＄％＾＄＆％＾（＆我“地锅一家人（家常炒菜）ａｂｃａ@￥@#%#ｓｄ🎈🎉ｆ我E２３４３４５んエォサ６３＃＄％＾＄＆％＾（＆我“",
	"1|1152921577701769238|家湘脆皮鸡(清扬路店)褚时健要退休了，依照自然规律，他会被淡忘，很难说一百年后，还会有多少人会去咀嚼他的故事。但是经营家族企业，谁不想做百年老店？ 这是褚氏企业必须面临的矛盾，这些年，褚橙的广告语一直是：“人生总有起落，精神终可传承”。当褚时健的光环消散后，褚橙是否还可以卖的那么贵，卖得那么好？ 褚时健管不了那么多了，那是下一代的事情。早在三年前，他就对《中国企业家》的记者说过：“我已经甘心了，我筋疲力尽了。”",
	"2|1152921628226355217|冰火公社麻辣烫",
	"3|1152921584588816406|杨铭宇黄\n鸡米饭(五星家园店)",
	"4|1152921538896068625|李克强从八个方面对过去一年的工作进行了回顾：创新和完善宏观调控，经济保持平稳运行；扎实打好三大攻坚战，重点任务取得积极进展；深化供给侧结构性改革，实体经济活力不断释放；深入实施创新驱动发展战略，创新能力和效率进一步提升；加大改革开放力度，发展动力继续增强；统筹城乡区域发展，良性互动格局加快形成；坚持在发展中保障和改善民生，改革发展成果更多更公平惠及人民群众；推进法治政府建设和治理创新，保持社会和谐稳定。",
	"5|3458764672814809215|淇帕创作（五星家园店）",
	"6|1152921672576925708|周宇记黄焖鸡米饭(金匱苑店)",
	"7|1152921708773769223|美丽家常菜馆",
	"8|1152921632588431382|美粥柒天（五星家园旗舰店）",
	"9|1152921711391014924|老北京炸酱面(五星家园店)",
	"10|1152921698518695953|三福酸菜鱼牛蛙店(清扬路店)",
	"11|3458764702137188504|吃货小店（五星家园店）",
	`12|1152921752692326422|漓江云姨桂林米粉（南下塘店）丁书苗为原铁道部部长刘志军身边的关键女商人，
号称高铁一姐。此番拍卖丁书苗资产的法院，即是负责丁书苗案的北京市二中院。早在2018年12月14日，
新京报独家报道，丁书苗旗下位于北京CBD的一处五星级的伯豪瑞廷大酒店，即将被北京市二中院拍卖，
评估价16亿元，起拍价11.27亿元。不过，12月25日，法院发布北京伯豪瑞廷酒店有限责任公司的100%股权变更公告，
"因出现法定事由，决定暂缓拍卖"。据介绍，北京伯豪瑞廷酒店于2008年7月开业，2009年被评定为五星级饭店。
是一家集客房、餐饮、会议、康乐功能为一体的商务酒店。酒店营业性建筑面积43000平方米。地处CBD核心商圈，
交通位置便利，服务功能齐全，获评“2015年度优秀会议酒店”，第十五届中国饭店金马奖。不过，
在2018年12月25日上午9时16分，法院发布北京伯豪瑞廷酒店有限责任公司的100%股权变更公告，
“因出现法定事由，决定暂缓拍卖”。据北京市二中院2月15日发布的通告，第一次拍卖被暂缓，
盖因法院接到多个电话实名举报，反映在拍卖过程中存在串标、围标、毁标等严重影响拍卖公平公正等行为，
经研究决定，二中院依法暂缓了该次拍卖，并就举报涉及的情况进行了相关的调查及处理。今年1月，
新京报报道，北京市第二中级人民法院将于2019年2月2日10时至2019年2月3日10时止（延时的除外）
进行公开拍卖活动，拍卖标的为北京伯豪瑞廷酒店有限责任公司的100%股权，评估价161064.566232万元，
起拍价144958.1096万元，保证金14000万元，增价幅度600万元。新京报记者注意到，重新开拍后，
该资产的起拍价提高了约3亿元。`,
	"13|1152921618071945228|吴记酸辣汤",
	"14|1152921524996145174|阿姨奶茶(清名一村店)",
	"15|3458764732797550744|墨冉小橱（家常炒菜）",
	"16|1152921734510018567|堡莱坞炸鸡汉堡（芦庄店）",
	"17|1152921594567065617|1+7水饺店(金科世界城店)",
	"18|1152921562266730504|豫味一品土豆粉（家乐福店）",
	"19|1152921553857150990|柒掌柜（五星家园店）陆慷表示，网上已经披露的有关部门提供的信息，重点介绍了加拿大公民康明凯涉嫌非法刺探、窃取中国国家秘密和情报的有关案件情况。根据有关方面披露的情况，案件侦办取得了重要进展，至于案件是否还涉及到其他人员、因素，包括你提到的，是否涉及中国公民，办案机关正在依法继续办理此案。能够披露的相关信息现已对外披露。",
	"20|1152921569279606801|山东水饺（茂业天地店）",
	"21|1152921556910604301|享吃私房披萨店",
	"22|1152921634152906774|食之味骨汤麻辣烫（通扬路店）",
	"23|1152921623264493590|腾宇记黄焖鸡米饭(瑞星家园店)",
	"24|1152921753204031511|胥先生鸭血粉丝（人民医院店）",
	"25|1152921708723437585|家常菜馆(中南路店)",
	`26|1152921596790046742|沙县小吃(五星人民医院店)<div class="dropdown global-dropdown">
	<button class="global-dropdown-toggle" data-toggle="dropdown" type="button">
	<span class="sr-only">Toggle navigation</span>
	<i aria-hidden='true' data-hidden = "true" class="fa fa-bars"></i>
	</button>
			<div class="dropdown-menu-nav global-dropdown-menu">
	<ul>
		<li class="home active">
<a title="Projects" class="dashboard-shortcuts-projects" href="/dashboard/projects">
			<div class="shortcut-mappings">
				<div class="key">
						<i aria-label="hidden" class="fa fa-arrow-up"></i>
						</div>
					</div>
				</a>
			</li>
		</ul>
		</div>
	</div>`,
	"27|1152921573658460169|小胖子功夫龙虾（无锡总店）",
	"28|1152921644617695244|小四川家常菜（扬名店）",
	"29|1152921546546479126|特色圆盅(前进路店)",
	"30|1152921509242339341|桥头排骨\n（金城路店）加拿大籍人员康明凯（Michael John Kovrig）窃取、刺探国家秘密和情报案侦办工作已取得重要进展。有关部门介绍，康明凯自2017年以来，经常持普通护照和商务签证入境，通过中国境内的关系人，窃取、刺探中国敏感信息和情报。迈克尔（Spavor Michael Peter Todd)是康明凯的重要情报关系人，向康明凯提供情报。",
}
