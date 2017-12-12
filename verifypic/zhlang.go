package verifypic

import "math/rand"
import "time"

func GenRandZh(leng int) string {
	zhmap := map[int]string{0: "龟", 1: "龙", 2: "龄", 3: "齿", 4: "齐", 5: "鼻", 6: "鼠", 7: "鼓", 8: "鼎", 9: "默", 10: "黔", 11: "黑", 12: "黎", 13: "黍", 14: "黄", 15: "麻", 16: "麸", 17: "麦", 18: "鹿", 19: "鹰", 20: "鹦", 21: "鹤", 22: "鹏", 23: "鹊", 24: "鹉", 25: "鹅", 26: "鹃", 27: "鸿", 28: "鸽", 29: "鸵", 30: "鸳", 31: "鸯", 32: "鸭", 33: "鸦", 34: "鸥", 35: "鸣", 36: "鸡", 37: "鸠", 38: "鸟", 39: "鳞", 40: "鳖", 41: "鳍", 42: "鳄", 43: "鲸", 44: "鲫", 45: "鲤", 46: "鲜", 47: "鲁", 48: "鱼", 49: "魔", 50: "魏", 51: "魄", 52: "魂", 53: "魁", 54: "鬼", 55: "鬓", 56: "高", 57: "髓", 58: "骨", 59: "骤", 60: "骡", 61: "骚", 62: "骗", 63: "骑", 64: "骏", 65: "验", 66: "骇", 67: "骆", 68: "骄", 69: "骂", 70: "驾", 71: "驼", 72: "驻", 73: "驹", 74: "驶", 75: "驴", 76: "驳", 77: "驱", 78: "驰", 79: "驯", 80: "驮", 81: "马", 82: "香", 83: "首", 84: "馒", 85: "馏", 86: "馍", 87: "馋", 88: "馆", 89: "馅", 90: "馁", 91: "饿", 92: "饼", 93: "饺", 94: "饶", 95: "饵", 96: "饲", 97: "饱", 98: "饰", 99: "饮", 100: "饭", 101: "饥", 102: "餐", 103: "食", 104: "飞", 105: "飘", 106: "飒", 107: "风", 108: "颤", 109: "颠", 110: "额", 111: "颜", 112: "题", 113: "颗", 114: "颖", 115: "颓", 116: "频", 117: "颊", 118: "颈", 119: "颇", 120: "领", 121: "颅", 122: "预", 123: "颂", 124: "颁", 125: "顿", 126: "顾", 127: "顽", 128: "须", 129: "顺", 130: "项", 131: "顷", 132: "顶", 133: "页", 134: "韵", 135: "音", 136: "韭", 137: "韩", 138: "韧", 139: "鞠", 140: "鞍", 141: "鞋", 142: "靶", 143: "靴", 144: "革", 145: "面", 146: "靡", 147: "靠", 148: "非", 149: "静", 150: "靖", 151: "青", 152: "霹", 153: "霸", 154: "露", 155: "霞", 156: "霜", 157: "霎", 158: "霍", 159: "霉", 160: "震", 161: "需", 162: "雾", 163: "雹", 164: "雷", 165: "零", 166: "雳", 167: "雪", 168: "雨", 169: "雕", 170: "雏", 171: "雌", 172: "雇", 173: "集", 174: "雅", 175: "雄", 176: "雁", 177: "雀", 178: "难", 179: "隶", 180: "隧", 181: "障", 182: "隙", 183: "隘", 184: "隔", 185: "隐", 186: "随", 187: "隆", 188: "隅", 189: "陷", 190: "陶", 191: "陵", 192: "陪", 193: "险", 194: "陨", 195: "除", 196: "院", 197: "陡", 198: "陕", 199: "限", 200: "降", 201: "陌", 202: "陋", 203: "陈", 204: "陆", 205: "际", 206: "附", 207: "阿", 208: "阻", 209: "阶", 210: "阵", 211: "阴", 212: "阳", 213: "防", 214: "阱", 215: "队", 216: "阔", 217: "阐", 218: "阎", 219: "阅", 220: "阁", 221: "阀", 222: "闽", 223: "闻", 224: "闺", 225: "闹", 226: "闸", 227: "闷", 228: "间", 229: "闲", 230: "闰", 231: "闯", 232: "问", 233: "闭", 234: "闪", 235: "门", 236: "长", 237: "镶", 238: "镰", 239: "镣", 240: "镜", 241: "镐", 242: "镊", 243: "镇", 244: "镀", 245: "锻", 246: "锹", 247: "锰", 248: "锯", 249: "键", 250: "锭", 251: "锨", 252: "锦", 253: "锥", 254: "锤", 255: "锣", 256: "锡", 257: "锚", 258: "错", 259: "锐", 260: "锌", 261: "锋", 262: "锉", 263: "锈", 264: "锅", 265: "锄", 266: "锁", 267: "销", 268: "链", 269: "铺", 270: "铸", 271: "银", 272: "铲", 273: "铭", 274: "铣", 275: "铡", 276: "铝", 277: "铜", 278: "铛", 279: "铐", 280: "铆", 281: "铅", 282: "铃", 283: "铁", 284: "钾", 285: "钻", 286: "钳", 287: "钱", 288: "钮", 289: "钩", 290: "钧", 291: "钦", 292: "钥", 293: "钢", 294: "钠", 295: "钟", 296: "钞", 297: "钝", 298: "钙", 299: "钓", 300: "钉", 301: "针", 302: "鉴", 303: "金", 304: "量", 305: "野", 306: "重", 307: "里", 308: "释", 309: "采", 310: "醒", 311: "醋", 312: "醉", 313: "醇", 314: "酿", 315: "酸", 316: "酷", 317: "酵", 318: "酱", 319: "酬", 320: "酪", 321: "酥", 322: "酣", 323: "酝", 324: "酗", 325: "酒", 326: "配", 327: "酌", 328: "鄙", 329: "都", 330: "郭", 331: "部", 332: "郑", 333: "郎", 334: "郊", 335: "郁", 336: "邻", 337: "邮", 338: "邪", 339: "邦", 340: "那", 341: "邢", 342: "邓", 343: "邑", 344: "邀", 345: "避", 346: "遵", 347: "遮", 348: "遭", 349: "遥", 350: "遣", 351: "遗", 352: "道", 353: "遏", 354: "遍", 355: "遇", 356: "遂", 357: "逾", 358: "逼", 359: "逻", 360: "逸", 361: "逮", 362: "逢", 363: "造", 364: "速", 365: "逞", 366: "逝", 367: "逛", 368: "通", 369: "逗", 370: "途", 371: "递", 372: "逐", 373: "透", 374: "逊", 375: "选", 376: "逆", 377: "逃", 378: "适", 379: "送", 380: "退", 381: "追", 382: "迹", 383: "迷", 384: "述", 385: "迫", 386: "迟", 387: "连", 388: "违", 389: "远", 390: "进", 391: "这", 392: "还", 393: "返", 394: "近", 395: "运", 396: "迎", 397: "迈", 398: "过", 399: "迅", 400: "迄", 401: "迂", 402: "迁", 403: "达", 404: "辽", 405: "边", 406: "辱", 407: "辰", 408: "辫", 409: "辩", 410: "辨", 411: "辣", 412: "辟", 413: "辞", 414: "辜", 415: "辛", 416: "辙", 417: "辖", 418: "辕", 419: "输", 420: "辑", 421: "辐", 422: "辉", 423: "辈", 424: "辆", 425: "辅", 426: "较", 427: "轿", 428: "载", 429: "轻", 430: "轴", 431: "轰", 432: "软", 433: "轮", 434: "转", 435: "轩", 436: "轨", 437: "轧", 438: "车", 439: "躺", 440: "躲", 441: "躯", 442: "躬", 443: "身", 444: "躏", 445: "躁", 446: "蹲", 447: "蹭", 448: "蹬", 449: "蹦", 450: "蹋", 451: "蹈", 452: "蹄", 453: "蹂", 454: "踱", 455: "踪", 456: "踩", 457: "踢", 458: "踏", 459: "踊", 460: "跺", 461: "跷", 462: "践", 463: "跳", 464: "路", 465: "跪", 466: "跨", 467: "跟", 468: "距", 469: "跛", 470: "跑", 471: "跌", 472: "跋", 473: "跃", 474: "趾", 475: "趴", 476: "足", 477: "趣", 478: "趟", 479: "趋", 480: "越", 481: "超", 482: "趁", 483: "起", 484: "赶", 485: "赵", 486: "赴", 487: "走", 488: "赫", 489: "赦", 490: "赤", 491: "赢", 492: "赡", 493: "赠", 494: "赞", 495: "赛", 496: "赚", 497: "赘", 498: "赖", 499: "赔", 500: "赐", 501: "赏", 502: "赎", 503: "赌", 504: "赋", 505: "赊", 506: "资", 507: "赃", 508: "赂", 509: "赁", 510: "贿", 511: "贾", 512: "贼", 513: "贺", 514: "费", 515: "贸", 516: "贷", 517: "贵", 518: "贴", 519: "贱", 520: "贰", 521: "贯", 522: "贮", 523: "购", 524: "贬", 525: "贫", 526: "贪", 527: "贩", 528: "质", 529: "货", 530: "账", 531: "败", 532: "贤", 533: "责", 534: "财", 535: "贡", 536: "负", 537: "贞", 538: "贝", 539: "貌", 540: "豺", 541: "豹", 542: "豫", 543: "豪", 544: "象", 545: "豌", 546: "豆", 547: "豁", 548: "谷", 549: "谴", 550: "谱", 551: "谭", 552: "谬", 553: "谨", 554: "谦", 555: "谤", 556: "谣", 557: "谢", 558: "谜", 559: "谚", 560: "谓", 561: "谒", 562: "谐", 563: "谎", 564: "谍", 565: "谋", 566: "谊", 567: "谈", 568: "谆", 569: "谅", 570: "调", 571: "谁", 572: "课", 573: "诽", 574: "读", 575: "诺", 576: "诸", 577: "请", 578: "诵", 579: "说", 580: "诲", 581: "诱", 582: "误", 583: "语", 584: "诬", 585: "诫", 586: "详", 587: "该", 588: "询", 589: "诡", 590: "诞", 591: "话", 592: "诚", 593: "诗", 594: "试", 595: "译", 596: "词", 597: "诊", 598: "诉", 599: "诈", 600: "识", 601: "诅", 602: "评", 603: "证", 604: "诀", 605: "访", 606: "设", 607: "讽", 608: "讼", 609: "论", 610: "讹", 611: "许", 612: "讶", 613: "讳", 614: "讲", 615: "记", 616: "讯", 617: "议", 618: "训", 619: "让", 620: "讨", 621: "讥", 622: "认", 623: "订", 624: "计", 625: "譬", 626: "警", 627: "誓", 628: "誊", 629: "誉", 630: "言", 631: "触", 632: "解", 633: "角", 634: "觉", 635: "览", 636: "视", 637: "觅", 638: "规", 639: "观", 640: "见", 641: "覆", 642: "要", 643: "西", 644: "襟", 645: "褪", 646: "褥", 647: "褒", 648: "褐", 649: "褂", 650: "裹", 651: "裸", 652: "裳", 653: "裤", 654: "裙", 655: "裕", 656: "裆", 657: "装", 658: "裂", 659: "裁", 660: "袱", 661: "袭", 662: "被", 663: "袜", 664: "袖", 665: "袒", 666: "袍", 667: "袋", 668: "袄", 669: "袁", 670: "衷", 671: "衰", 672: "衬", 673: "衫", 674: "衩", 675: "表", 676: "补", 677: "衣", 678: "衡", 679: "衙", 680: "街", 681: "衔", 682: "衍", 683: "行", 684: "衅", 685: "血", 686: "蠢", 687: "蠕", 688: "蟹", 689: "蟥", 690: "蟋", 691: "蟆", 692: "蟀", 693: "螺", 694: "螟", 695: "融", 696: "螃", 697: "蝶", 698: "蝴", 699: "蝠", 700: "蝙", 701: "蝗", 702: "蝎", 703: "蝌", 704: "蝉", 705: "蝇", 706: "蜻", 707: "蜡", 708: "蜜", 709: "蜘", 710: "蜗", 711: "蜕", 712: "蜓", 713: "蜒", 714: "蜈", 715: "蜂", 716: "蜀", 717: "蛾", 718: "蛹", 719: "蛮", 720: "蛤", 721: "蛛", 722: "蛙", 723: "蛔", 724: "蛋", 725: "蛉", 726: "蛇", 727: "蛆", 728: "蛀", 729: "蚯", 730: "蚪", 731: "蚤", 732: "蚣", 733: "蚜", 734: "蚕", 735: "蚓", 736: "蚌", 737: "蚊", 738: "蚂", 739: "蚁", 740: "蚀", 741: "虾", 742: "虽", 743: "虹", 744: "虱", 745: "虫", 746: "虚", 747: "虑", 748: "虐", 749: "虏", 750: "虎", 751: "蘸", 752: "蘑", 753: "藻", 754: "藤", 755: "藕", 756: "藐", 757: "藏", 758: "薯", 759: "薪", 760: "薛", 761: "薇", 762: "薄", 763: "蕾", 764: "蕴", 765: "蕊", 766: "蕉", 767: "蔽", 768: "蔼", 769: "蔬", 770: "蔫", 771: "蔚", 772: "蔗", 773: "蔓", 774: "蔑", 775: "蓬", 776: "蓝", 777: "蓖", 778: "蓉", 779: "蓄", 780: "蒿", 781: "蒸", 782: "蒲", 783: "蒜", 784: "蒙", 785: "蒋", 786: "蒂", 787: "葵", 788: "葱", 789: "葬", 790: "葫", 791: "董", 792: "葡", 793: "葛", 794: "著", 795: "落", 796: "萨", 797: "萧", 798: "营", 799: "萤", 800: "萝", 801: "萎", 802: "萍", 803: "萌", 804: "萄", 805: "菲", 806: "菱", 807: "菩", 808: "菠", 809: "菜", 810: "菌", 811: "菊", 812: "菇", 813: "莽", 814: "莺", 815: "莹", 816: "获", 817: "莲", 818: "莱", 819: "莫", 820: "莉", 821: "荸", 822: "荷", 823: "药", 824: "荧", 825: "荤", 826: "荣", 827: "荡", 828: "荠", 829: "荞", 830: "荚", 831: "荔", 832: "荒", 833: "荐", 834: "草", 835: "荆", 836: "茸", 837: "茶", 838: "茵", 839: "茴", 840: "茬", 841: "茫", 842: "茧", 843: "茎", 844: "茉", 845: "茅", 846: "茄", 847: "范", 848: "茂", 849: "茁", 850: "苹", 851: "英", 852: "苫", 853: "苦", 854: "若", 855: "苟", 856: "苞", 857: "苛", 858: "苗", 859: "苔", 860: "苏", 861: "苍", 862: "苇", 863: "芽", 864: "芹", 865: "芳", 866: "花", 867: "芯", 868: "芭", 869: "芬", 870: "芦", 871: "芥", 872: "芝", 873: "芜", 874: "芙", 875: "芒", 876: "芍", 877: "芋", 878: "节", 879: "艾", 880: "艺", 881: "艳", 882: "色", 883: "艰", 884: "良", 885: "艘", 886: "艇", 887: "船", 888: "舷", 889: "舶", 890: "舵", 891: "舱", 892: "舰", 893: "般", 894: "航", 895: "舟", 896: "舞", 897: "舔", 898: "舒", 899: "舍", 900: "舌", 901: "舆", 902: "舅", 903: "舀", 904: "臼", 905: "致", 906: "至", 907: "臭", 908: "自", 909: "臣", 910: "臊", 911: "臂", 912: "臀", 913: "膳", 914: "膨", 915: "膝", 916: "膜", 917: "膛", 918: "膘", 919: "膏", 920: "膊", 921: "膀", 922: "腿", 923: "腾", 924: "腻", 925: "腺", 926: "腹", 927: "腰", 928: "腮", 929: "腥", 930: "腕", 931: "腔", 932: "腐", 933: "腌", 934: "腋", 935: "腊", 936: "脾", 937: "脸", 938: "脱", 939: "脯", 940: "脚", 941: "脖", 942: "脓", 943: "脑", 944: "脐", 945: "脏", 946: "脊", 947: "脉", 948: "脆", 949: "脂", 950: "能", 951: "胸", 952: "胶", 953: "胳", 954: "胰", 955: "胯", 956: "胧", 957: "胡", 958: "胞", 959: "胜", 960: "胚", 961: "胖", 962: "胎", 963: "背", 964: "胆", 965: "胃", 966: "胁", 967: "胀", 968: "肿", 969: "肾", 970: "肺", 971: "肴", 972: "育", 973: "肯", 974: "肮", 975: "肪", 976: "肩", 977: "肥", 978: "肤", 979: "肢", 980: "股", 981: "肠", 982: "肝", 983: "肛", 984: "肚", 985: "肘", 986: "肖", 987: "肌", 988: "肋", 989: "肉", 990: "肆", 991: "肄", 992: "肃", 993: "聪", 994: "聚", 995: "聘", 996: "联", 997: "职", 998: "聋", 999: "聊", 1000: "聂", 1001: "耿", 1002: "耽", 1003: "耻", 1004: "耸", 1005: "耳", 1006: "耙", 1007: "耘", 1008: "耗", 1009: "耕", 1010: "耐", 1011: "耍", 1012: "而", 1013: "者", 1014: "考", 1015: "老", 1016: "耀", 1017: "翼", 1018: "翻", 1019: "翰", 1020: "翩", 1021: "翠", 1022: "翘", 1023: "翔", 1024: "翎", 1025: "翅", 1026: "翁", 1027: "羽", 1028: "羹", 1029: "群", 1030: "羡", 1031: "羞", 1032: "羔", 1033: "美", 1034: "羊", 1035: "署", 1036: "置", 1037: "罪", 1038: "罩", 1039: "罢", 1040: "罚", 1041: "罗", 1042: "罕", 1043: "网", 1044: "罐", 1045: "缺", 1046: "缸", 1047: "缴", 1048: "缰", 1049: "缭", 1050: "缩", 1051: "缨", 1052: "缤", 1053: "缠", 1054: "缝", 1055: "缚", 1056: "缘", 1057: "编", 1058: "缕", 1059: "缔", 1060: "缓", 1061: "缎", 1062: "缆", 1063: "缅", 1064: "缀", 1065: "绿", 1066: "绽", 1067: "综", 1068: "绸", 1069: "绷", 1070: "绵", 1071: "维", 1072: "绳", 1073: "绰", 1074: "续", 1075: "绪", 1076: "绩", 1077: "继", 1078: "绣", 1079: "绢", 1080: "统", 1081: "绞", 1082: "绝", 1083: "络", 1084: "给", 1085: "绘", 1086: "绕", 1087: "结", 1088: "绒", 1089: "绑", 1090: "经", 1091: "绎", 1092: "绍", 1093: "绊", 1094: "终", 1095: "织", 1096: "细", 1097: "绅", 1098: "组", 1099: "练", 1100: "线", 1101: "纽", 1102: "纺", 1103: "纹", 1104: "纸", 1105: "纷", 1106: "纵", 1107: "纳", 1108: "纲", 1109: "纱", 1110: "纯", 1111: "纬", 1112: "纫", 1113: "纪", 1114: "级", 1115: "约", 1116: "纤", 1117: "红", 1118: "纠", 1119: "繁", 1120: "絮", 1121: "累", 1122: "紫", 1123: "紧", 1124: "索", 1125: "素", 1126: "紊", 1127: "系", 1128: "糯", 1129: "糠", 1130: "糟", 1131: "糜", 1132: "糙", 1133: "糖", 1134: "糕", 1135: "糊", 1136: "精", 1137: "粹", 1138: "粱", 1139: "粮", 1140: "粪", 1141: "粥", 1142: "粤", 1143: "粟", 1144: "粘", 1145: "粗", 1146: "粒", 1147: "粉", 1148: "籽", 1149: "类", 1150: "米", 1151: "籍", 1152: "簿", 1153: "簸", 1154: "簇", 1155: "篷", 1156: "篱", 1157: "篮", 1158: "篡", 1159: "篙", 1160: "篓", 1161: "篇", 1162: "箱", 1163: "箭", 1164: "箫", 1165: "箩", 1166: "管", 1167: "算", 1168: "箕", 1169: "箍", 1170: "简", 1171: "签", 1172: "筹", 1173: "筷", 1174: "筝", 1175: "筛", 1176: "策", 1177: "答", 1178: "筒", 1179: "筑", 1180: "筐", 1181: "筏", 1182: "筋", 1183: "等", 1184: "笼", 1185: "第", 1186: "笨", 1187: "符", 1188: "笤", 1189: "笛", 1190: "笙", 1191: "笔", 1192: "笑", 1193: "笋", 1194: "笆", 1195: "竿", 1196: "竹", 1197: "端", 1198: "竭", 1199: "童", 1200: "竣", 1201: "章", 1202: "竟", 1203: "竞", 1204: "站", 1205: "竖", 1206: "立", 1207: "窿", 1208: "窥", 1209: "窟", 1210: "窝", 1211: "窜", 1212: "窘", 1213: "窗", 1214: "窖", 1215: "窒", 1216: "窑", 1217: "窍", 1218: "窄", 1219: "窃", 1220: "突", 1221: "穿", 1222: "空", 1223: "穷", 1224: "究", 1225: "穴", 1226: "穗", 1227: "穆", 1228: "稿", 1229: "稽", 1230: "稼", 1231: "稻", 1232: "稳", 1233: "稠", 1234: "稚", 1235: "税", 1236: "稍", 1237: "程", 1238: "稀", 1239: "秽", 1240: "移", 1241: "秸", 1242: "称", 1243: "积", 1244: "秫", 1245: "秩", 1246: "秧", 1247: "秦", 1248: "秤", 1249: "租", 1250: "秘", 1251: "秕", 1252: "秒", 1253: "科", 1254: "种", 1255: "秋", 1256: "秉", 1257: "秆", 1258: "秃", 1259: "私", 1260: "秀", 1261: "禾", 1262: "禽", 1263: "离", 1264: "福", 1265: "禁", 1266: "禀", 1267: "祸", 1268: "祷", 1269: "祭", 1270: "票", 1271: "祥", 1272: "祠", 1273: "祟", 1274: "神", 1275: "祝", 1276: "祖", 1277: "祈", 1278: "社", 1279: "礼", 1280: "示", 1281: "礁", 1282: "磷", 1283: "磨", 1284: "磕", 1285: "磅", 1286: "磁", 1287: "碾", 1288: "碴", 1289: "碳", 1290: "碱", 1291: "碰", 1292: "碧", 1293: "碟", 1294: "碘", 1295: "碗", 1296: "碑", 1297: "碎", 1298: "碍", 1299: "碌", 1300: "碉", 1301: "硼", 1302: "确", 1303: "硬", 1304: "硫", 1305: "硝", 1306: "硕", 1307: "硅", 1308: "础", 1309: "砾", 1310: "砸", 1311: "破", 1312: "砰", 1313: "砚", 1314: "砖", 1315: "研", 1316: "砍", 1317: "砌", 1318: "砂", 1319: "码", 1320: "矿", 1321: "矾", 1322: "石", 1323: "矮", 1324: "短", 1325: "矫", 1326: "矩", 1327: "知", 1328: "矢", 1329: "矛", 1330: "矗", 1331: "瞻", 1332: "瞳", 1333: "瞭", 1334: "瞬", 1335: "瞪", 1336: "瞧", 1337: "瞒", 1338: "瞎", 1339: "瞄", 1340: "睹", 1341: "睬", 1342: "睦", 1343: "督", 1344: "睡", 1345: "睛", 1346: "睁", 1347: "着", 1348: "眼", 1349: "眷", 1350: "眶", 1351: "眯", 1352: "眨", 1353: "眠", 1354: "真", 1355: "看", 1356: "眉", 1357: "省", 1358: "盾", 1359: "盼", 1360: "盹", 1361: "相", 1362: "直", 1363: "盲", 1364: "盯", 1365: "目", 1366: "盟", 1367: "盛", 1368: "盘", 1369: "盗", 1370: "盖", 1371: "盔", 1372: "盒", 1373: "监", 1374: "盐", 1375: "盏", 1376: "益", 1377: "盈", 1378: "盆", 1379: "盅", 1380: "皿", 1381: "皱", 1382: "皮", 1383: "皇", 1384: "皆", 1385: "的", 1386: "皂", 1387: "百", 1388: "白", 1389: "登", 1390: "癣", 1391: "癞", 1392: "癌", 1393: "瘾", 1394: "瘸", 1395: "瘫", 1396: "瘪", 1397: "瘩", 1398: "瘦", 1399: "瘤", 1400: "瘟", 1401: "痹", 1402: "痴", 1403: "痰", 1404: "痪", 1405: "痢", 1406: "痛", 1407: "痘", 1408: "痕", 1409: "痒", 1410: "痊", 1411: "症", 1412: "病", 1413: "疾", 1414: "疼", 1415: "疹", 1416: "疲", 1417: "疯", 1418: "疮", 1419: "疫", 1420: "疤", 1421: "疟", 1422: "疚", 1423: "疙", 1424: "疗", 1425: "疑", 1426: "疏", 1427: "疆", 1428: "畸", 1429: "畴", 1430: "番", 1431: "畦", 1432: "略", 1433: "畜", 1434: "留", 1435: "畔", 1436: "畏", 1437: "界", 1438: "畅", 1439: "画", 1440: "甸", 1441: "男", 1442: "电", 1443: "申", 1444: "甲", 1445: "由", 1446: "田", 1447: "甫", 1448: "甩", 1449: "用", 1450: "甥", 1451: "生", 1452: "甜", 1453: "甚", 1454: "甘", 1455: "瓷", 1456: "瓶", 1457: "瓮", 1458: "瓦", 1459: "瓤", 1460: "瓣", 1461: "瓢", 1462: "瓜", 1463: "璧", 1464: "璃", 1465: "瑰", 1466: "瑟", 1467: "瑞", 1468: "琼", 1469: "琴", 1470: "琳", 1471: "琢", 1472: "琐", 1473: "琉", 1474: "理", 1475: "琅", 1476: "球", 1477: "班", 1478: "珠", 1479: "珍", 1480: "珊", 1481: "玻", 1482: "玷", 1483: "玲", 1484: "现", 1485: "环", 1486: "玫", 1487: "玩", 1488: "玛", 1489: "玖", 1490: "王", 1491: "玉", 1492: "率", 1493: "玄", 1494: "猿", 1495: "猾", 1496: "猴", 1497: "献", 1498: "猬", 1499: "猫", 1500: "猪", 1501: "猩", 1502: "猜", 1503: "猛", 1504: "猖", 1505: "猎", 1506: "狼", 1507: "狸", 1508: "狱", 1509: "狰", 1510: "狮", 1511: "狭", 1512: "独", 1513: "狡", 1514: "狠", 1515: "狞", 1516: "狗", 1517: "狐", 1518: "狈", 1519: "狂", 1520: "犹", 1521: "状", 1522: "犯", 1523: "犬", 1524: "犁", 1525: "犀", 1526: "牺", 1527: "特", 1528: "牵", 1529: "牲", 1530: "物", 1531: "牧", 1532: "牢", 1533: "牡", 1534: "牛", 1535: "牙", 1536: "牍", 1537: "牌", 1538: "版", 1539: "片", 1540: "爽", 1541: "爹", 1542: "爸", 1543: "爷", 1544: "父", 1545: "爵", 1546: "爱", 1547: "爬", 1548: "爪", 1549: "爆", 1550: "燥", 1551: "燕", 1552: "燎", 1553: "燃", 1554: "熬", 1555: "熟", 1556: "熙", 1557: "熔", 1558: "熏", 1559: "熊", 1560: "熄", 1561: "煮", 1562: "照", 1563: "煤", 1564: "煞", 1565: "煎", 1566: "煌", 1567: "然", 1568: "焰", 1569: "焦", 1570: "焚", 1571: "焙", 1572: "焕", 1573: "焊", 1574: "烹", 1575: "热", 1576: "烫", 1577: "烧", 1578: "烦", 1579: "烤", 1580: "烟", 1581: "烛", 1582: "烙", 1583: "烘", 1584: "烈", 1585: "烂", 1586: "烁", 1587: "炼", 1588: "点", 1589: "炸", 1590: "炮", 1591: "炭", 1592: "炬", 1593: "炫", 1594: "炕", 1595: "炒", 1596: "炎", 1597: "炊", 1598: "炉", 1599: "灿", 1600: "灾", 1601: "灼", 1602: "灸", 1603: "灶", 1604: "灵", 1605: "灰", 1606: "灯", 1607: "灭", 1608: "火", 1609: "灌", 1610: "瀑", 1611: "濒", 1612: "激", 1613: "澳", 1614: "澡", 1615: "澜", 1616: "澎", 1617: "澈", 1618: "澄", 1619: "潮", 1620: "潭", 1621: "潦", 1622: "潜", 1623: "潘", 1624: "漾", 1625: "漱", 1626: "漫", 1627: "漩", 1628: "漠", 1629: "演", 1630: "漓", 1631: "漏", 1632: "漆", 1633: "漂", 1634: "滴", 1635: "滩", 1636: "滨", 1637: "滥", 1638: "滤", 1639: "满", 1640: "滞", 1641: "滚", 1642: "滔", 1643: "滓", 1644: "滑", 1645: "滋", 1646: "溺", 1647: "溶", 1648: "溯", 1649: "溪", 1650: "溢", 1651: "溜", 1652: "源", 1653: "溉", 1654: "溅", 1655: "溃", 1656: "湿", 1657: "湾", 1658: "湘", 1659: "湖", 1660: "湃", 1661: "渺", 1662: "游", 1663: "渴", 1664: "港", 1665: "温", 1666: "渤", 1667: "渣", 1668: "渡", 1669: "渠", 1670: "渗", 1671: "渔", 1672: "渐", 1673: "渊", 1674: "清", 1675: "添", 1676: "淹", 1677: "混", 1678: "淳", 1679: "深", 1680: "淮", 1681: "淫", 1682: "淤", 1683: "淡", 1684: "淘", 1685: "淑", 1686: "淌", 1687: "淋", 1688: "淆", 1689: "淀", 1690: "涵", 1691: "液", 1692: "涯", 1693: "涮", 1694: "涩", 1695: "涨", 1696: "涧", 1697: "润", 1698: "涤", 1699: "涣", 1700: "涡", 1701: "涝", 1702: "涛", 1703: "涕", 1704: "涎", 1705: "涌", 1706: "涉", 1707: "消", 1708: "涂", 1709: "浸", 1710: "海", 1711: "浴", 1712: "浮", 1713: "浪", 1714: "浩", 1715: "浦", 1716: "浙", 1717: "浓", 1718: "浑", 1719: "济", 1720: "测", 1721: "浊", 1722: "浇", 1723: "浆", 1724: "浅", 1725: "流", 1726: "派", 1727: "洽", 1728: "洼", 1729: "活", 1730: "洲", 1731: "洪", 1732: "津", 1733: "洞", 1734: "洛", 1735: "洗", 1736: "洒", 1737: "洋", 1738: "洁", 1739: "泽", 1740: "泼", 1741: "泻", 1742: "泵", 1743: "泳", 1744: "泰", 1745: "泪", 1746: "注", 1747: "泥", 1748: "泣", 1749: "波", 1750: "泡", 1751: "泞", 1752: "泛", 1753: "法", 1754: "泌", 1755: "泊", 1756: "泉", 1757: "泄", 1758: "沿", 1759: "沾", 1760: "沽", 1761: "沼", 1762: "治", 1763: "油", 1764: "沸", 1765: "河", 1766: "沮", 1767: "沫", 1768: "沪", 1769: "沧", 1770: "沦", 1771: "沥", 1772: "没", 1773: "沟", 1774: "沛", 1775: "沙", 1776: "沐", 1777: "沉", 1778: "沈", 1779: "沃", 1780: "汽", 1781: "汹", 1782: "汰", 1783: "汪", 1784: "汤", 1785: "污", 1786: "池", 1787: "江", 1788: "汞", 1789: "汛", 1790: "汗", 1791: "汉", 1792: "汇", 1793: "求", 1794: "汁", 1795: "永", 1796: "水", 1797: "氯", 1798: "氮", 1799: "氨", 1800: "氧", 1801: "氢", 1802: "氛", 1803: "气", 1804: "氓", 1805: "民", 1806: "氏", 1807: "毯", 1808: "毫", 1809: "毡", 1810: "毛", 1811: "毙", 1812: "毕", 1813: "比", 1814: "毒", 1815: "每", 1816: "母", 1817: "毅", 1818: "毁", 1819: "殿", 1820: "殷", 1821: "段", 1822: "殴", 1823: "殖", 1824: "残", 1825: "殊", 1826: "殉", 1827: "殃", 1828: "歼", 1829: "死", 1830: "歹", 1831: "歪", 1832: "歧", 1833: "武", 1834: "步", 1835: "此", 1836: "正", 1837: "止", 1838: "歌", 1839: "歉", 1840: "歇", 1841: "款", 1842: "欺", 1843: "欲", 1844: "欧", 1845: "欣", 1846: "欢", 1847: "次", 1848: "欠", 1849: "檬", 1850: "檩", 1851: "檐", 1852: "檀", 1853: "橱", 1854: "橡", 1855: "橙", 1856: "橘", 1857: "橄", 1858: "樱", 1859: "横", 1860: "模", 1861: "樟", 1862: "樊", 1863: "槽", 1864: "槐", 1865: "榴", 1866: "榨", 1867: "榜", 1868: "榛", 1869: "榕", 1870: "榔", 1871: "榆", 1872: "榄", 1873: "概", 1874: "楼", 1875: "楷", 1876: "楣", 1877: "楞", 1878: "楚", 1879: "楔", 1880: "椿", 1881: "椰", 1882: "椭", 1883: "椒", 1884: "椎", 1885: "植", 1886: "椅", 1887: "棺", 1888: "棵", 1889: "棱", 1890: "森", 1891: "棠", 1892: "棚", 1893: "棘", 1894: "棕", 1895: "棒", 1896: "棍", 1897: "棋", 1898: "棉", 1899: "检", 1900: "梳", 1901: "械", 1902: "梯", 1903: "梭", 1904: "梨", 1905: "梧", 1906: "梦", 1907: "梢", 1908: "梗", 1909: "梆", 1910: "梅", 1911: "梁", 1912: "桶", 1913: "桩", 1914: "桨", 1915: "桦", 1916: "桥", 1917: "档", 1918: "桑", 1919: "桐", 1920: "桌", 1921: "案", 1922: "框", 1923: "桅", 1924: "桃", 1925: "桂", 1926: "栽", 1927: "格", 1928: "根", 1929: "核", 1930: "样", 1931: "株", 1932: "校", 1933: "栗", 1934: "栖", 1935: "栓", 1936: "树", 1937: "栏", 1938: "栋", 1939: "栈", 1940: "标", 1941: "栅", 1942: "柿", 1943: "柴", 1944: "柳", 1945: "柱", 1946: "柬", 1947: "查", 1948: "柠", 1949: "柜", 1950: "柔", 1951: "染", 1952: "柒", 1953: "柑", 1954: "某", 1955: "柏", 1956: "柄", 1957: "枷", 1958: "架", 1959: "枯", 1960: "枫", 1961: "枪", 1962: "枣", 1963: "枢", 1964: "枝", 1965: "果", 1966: "枚", 1967: "林", 1968: "枕", 1969: "析", 1970: "枉", 1971: "构", 1972: "极", 1973: "板", 1974: "松", 1975: "杰", 1976: "杯", 1977: "杭", 1978: "杨", 1979: "来", 1980: "条", 1981: "杠", 1982: "束", 1983: "杜", 1984: "杖", 1985: "村", 1986: "材", 1987: "杏", 1988: "李", 1989: "杉", 1990: "杈", 1991: "杆", 1992: "权", 1993: "杂", 1994: "杀", 1995: "朽", 1996: "机", 1997: "朵", 1998: "朴", 1999: "朱", 2000: "术", 2001: "本", 2002: "末", 2003: "未", 2004: "木", 2005: "朦", 2006: "期", 2007: "朝", 2008: "望", 2009: "朗", 2010: "服", 2011: "朋", 2012: "有", 2013: "月", 2014: "最", 2015: "替", 2016: "曾", 2017: "曼", 2018: "曹", 2019: "更", 2020: "曲", 2021: "曙", 2022: "暴", 2023: "暮", 2024: "暗", 2025: "暖", 2026: "暑", 2027: "暇", 2028: "暂", 2029: "晾", 2030: "智", 2031: "晶", 2032: "晴", 2033: "晰", 2034: "景", 2035: "普", 2036: "晨", 2037: "晦", 2038: "晤", 2039: "晚", 2040: "晕", 2041: "晓", 2042: "晒", 2043: "晌", 2044: "晋", 2045: "晃", 2046: "显", 2047: "昼", 2048: "昵", 2049: "是", 2050: "昭", 2051: "昨", 2052: "昧", 2053: "春", 2054: "映", 2055: "星", 2056: "昙", 2057: "昔", 2058: "易", 2059: "昏", 2060: "明", 2061: "昌", 2062: "昆", 2063: "昂", 2064: "旺", 2065: "旷", 2066: "时", 2067: "旱", 2068: "旭", 2069: "旬", 2070: "早", 2071: "旨", 2072: "旧", 2073: "旦", 2074: "日", 2075: "既", 2076: "无", 2077: "旗", 2078: "族", 2079: "旋", 2080: "旅", 2081: "旁", 2082: "施", 2083: "方", 2084: "新", 2085: "斯", 2086: "断", 2087: "斩", 2088: "斧", 2089: "斥", 2090: "斤", 2091: "斟", 2092: "斜", 2093: "料", 2094: "斗", 2095: "斑", 2096: "斋", 2097: "文", 2098: "敷", 2099: "整", 2100: "敲", 2101: "数", 2102: "敬", 2103: "敦", 2104: "散", 2105: "敢", 2106: "敞", 2107: "敛", 2108: "教", 2109: "救", 2110: "敏", 2111: "敌", 2112: "效", 2113: "故", 2114: "政", 2115: "放", 2116: "攻", 2117: "改", 2118: "收", 2119: "支", 2120: "攘", 2121: "攒", 2122: "攀", 2123: "擦", 2124: "擒", 2125: "擎", 2126: "操", 2127: "擅", 2128: "擂", 2129: "撼", 2130: "撵", 2131: "撰", 2132: "撮", 2133: "播", 2134: "撬", 2135: "撩", 2136: "撤", 2137: "撞", 2138: "撕", 2139: "撒", 2140: "撑", 2141: "撇", 2142: "摹", 2143: "摸", 2144: "摩", 2145: "摧", 2146: "摘", 2147: "摔", 2148: "摊", 2149: "摇", 2150: "摆", 2151: "摄", 2152: "携", 2153: "搭", 2154: "搬", 2155: "搪", 2156: "搞", 2157: "搜", 2158: "搔", 2159: "搓", 2160: "搏", 2161: "搅", 2162: "搂", 2163: "搁", 2164: "搀", 2165: "揽", 2166: "援", 2167: "揭", 2168: "揪", 2169: "揩", 2170: "揣", 2171: "握", 2172: "揖", 2173: "插", 2174: "提", 2175: "描", 2176: "揍", 2177: "揉", 2178: "掺", 2179: "掸", 2180: "掷", 2181: "掰", 2182: "措", 2183: "掩", 2184: "推", 2185: "控", 2186: "接", 2187: "探", 2188: "掠", 2189: "掘", 2190: "掖", 2191: "排", 2192: "掐", 2193: "掏", 2194: "掌", 2195: "掉", 2196: "授", 2197: "掂", 2198: "掀", 2199: "捻", 2200: "捺", 2201: "捷", 2202: "捶", 2203: "据", 2204: "捧", 2205: "捣", 2206: "换", 2207: "捡", 2208: "损", 2209: "捞", 2210: "捕", 2211: "捐", 2212: "捏", 2213: "捎", 2214: "捍", 2215: "捌", 2216: "捉", 2217: "捆", 2218: "捅", 2219: "捂", 2220: "挽", 2221: "挺", 2222: "振", 2223: "挫", 2224: "挪", 2225: "挨", 2226: "挥", 2227: "挤", 2228: "挣", 2229: "挡", 2230: "挠", 2231: "挟", 2232: "挚", 2233: "挖", 2234: "挑", 2235: "挎", 2236: "按", 2237: "指", 2238: "挂", 2239: "持", 2240: "拿", 2241: "拾", 2242: "拼", 2243: "拷", 2244: "拴", 2245: "拳", 2246: "拱", 2247: "拯", 2248: "拭", 2249: "括", 2250: "择", 2251: "拨", 2252: "拧", 2253: "拦", 2254: "拥", 2255: "拣", 2256: "拢", 2257: "拟", 2258: "拜", 2259: "招", 2260: "拙", 2261: "拘", 2262: "拗", 2263: "拖", 2264: "拔", 2265: "拓", 2266: "拒", 2267: "拐", 2268: "拍", 2269: "拌", 2270: "拉", 2271: "拇", 2272: "拆", 2273: "担", 2274: "拄", 2275: "拂", 2276: "抽", 2277: "押", 2278: "抹", 2279: "抵", 2280: "抱", 2281: "抬", 2282: "披", 2283: "报", 2284: "护", 2285: "抢", 2286: "抡", 2287: "抠", 2288: "抛", 2289: "抚", 2290: "折", 2291: "抗", 2292: "抖", 2293: "投", 2294: "抓", 2295: "抒", 2296: "抑", 2297: "把", 2298: "抄", 2299: "技", 2300: "承", 2301: "找", 2302: "扼", 2303: "批", 2304: "扶", 2305: "扳", 2306: "扰", 2307: "扯", 2308: "扮", 2309: "扭", 2310: "扬", 2311: "扫", 2312: "扩", 2313: "执", 2314: "扣", 2315: "扛", 2316: "托", 2317: "扔", 2318: "打", 2319: "扒", 2320: "扑", 2321: "扎", 2322: "才", 2323: "手", 2324: "扇", 2325: "扁", 2326: "所", 2327: "房", 2328: "户", 2329: "戴", 2330: "戳", 2331: "截", 2332: "戚", 2333: "战", 2334: "或", 2335: "戒", 2336: "我", 2337: "成", 2338: "戏", 2339: "戈", 2340: "懦", 2341: "懒", 2342: "懊", 2343: "懈", 2344: "懂", 2345: "憾", 2346: "憨", 2347: "憔", 2348: "憎", 2349: "憋", 2350: "慷", 2351: "慰", 2352: "慨", 2353: "慧", 2354: "慢", 2355: "慕", 2356: "慎", 2357: "慌", 2358: "慈", 2359: "愿", 2360: "愧", 2361: "愤", 2362: "感", 2363: "愚", 2364: "愕", 2365: "意", 2366: "愉", 2367: "愈", 2368: "愁", 2369: "惹", 2370: "惶", 2371: "想", 2372: "惰", 2373: "惯", 2374: "惭", 2375: "惫", 2376: "惩", 2377: "惨", 2378: "惧", 2379: "惦", 2380: "惠", 2381: "惜", 2382: "惕", 2383: "惑", 2384: "惋", 2385: "惊", 2386: "情", 2387: "悼", 2388: "悴", 2389: "悲", 2390: "悯", 2391: "悬", 2392: "您", 2393: "悦", 2394: "患", 2395: "悠", 2396: "悟", 2397: "悔", 2398: "悍", 2399: "悉", 2400: "悄", 2401: "恼", 2402: "恶", 2403: "恳", 2404: "恰", 2405: "息", 2406: "恭", 2407: "恬", 2408: "恩", 2409: "恨", 2410: "恤", 2411: "恢", 2412: "恕", 2413: "恒", 2414: "恐", 2415: "恍", 2416: "恋", 2417: "恃", 2418: "总", 2419: "怯", 2420: "怪", 2421: "怨", 2422: "性", 2423: "急", 2424: "怠", 2425: "思", 2426: "怜", 2427: "怖", 2428: "怕", 2429: "怔", 2430: "怒", 2431: "怎", 2432: "态", 2433: "怀", 2434: "忿", 2435: "忽", 2436: "念", 2437: "忱", 2438: "快", 2439: "忧", 2440: "忠", 2441: "忙", 2442: "忘", 2443: "志", 2444: "忍", 2445: "忌", 2446: "忆", 2447: "必", 2448: "心", 2449: "徽", 2450: "德", 2451: "微", 2452: "循", 2453: "御", 2454: "徙", 2455: "徘", 2456: "得", 2457: "徒", 2458: "徐", 2459: "律", 2460: "徊", 2461: "很", 2462: "待", 2463: "径", 2464: "征", 2465: "往", 2466: "彼", 2467: "彻", 2468: "役", 2469: "影", 2470: "彰", 2471: "彭", 2472: "彬", 2473: "彪", 2474: "彩", 2475: "彤", 2476: "形", 2477: "录", 2478: "当", 2479: "归", 2480: "强", 2481: "弹", 2482: "弱", 2483: "弯", 2484: "弧", 2485: "弦", 2486: "弥", 2487: "张", 2488: "弟", 2489: "弛", 2490: "引", 2491: "弓", 2492: "式", 2493: "弊", 2494: "弄", 2495: "弃", 2496: "异", 2497: "开", 2498: "建", 2499: "廷", 2500: "延", 2501: "廓", 2502: "廊", 2503: "廉", 2504: "庸", 2505: "康", 2506: "庶", 2507: "庵", 2508: "庭", 2509: "座", 2510: "度", 2511: "废", 2512: "庞", 2513: "府", 2514: "庙", 2515: "店", 2516: "底", 2517: "应", 2518: "库", 2519: "庐", 2520: "序", 2521: "床", 2522: "庇", 2523: "庆", 2524: "庄", 2525: "广", 2526: "幽", 2527: "幼", 2528: "幻", 2529: "幸", 2530: "并", 2531: "年", 2532: "平", 2533: "干", 2534: "幢", 2535: "幕", 2536: "幔", 2537: "幌", 2538: "幅", 2539: "帽", 2540: "常", 2541: "帮", 2542: "席", 2543: "带", 2544: "帝", 2545: "帜", 2546: "帚", 2547: "帘", 2548: "帖", 2549: "帕", 2550: "帐", 2551: "希", 2552: "师", 2553: "帆", 2554: "帅", 2555: "布", 2556: "市", 2557: "币", 2558: "巾", 2559: "巷", 2560: "巴", 2561: "已", 2562: "己", 2563: "差", 2564: "巫", 2565: "巩", 2566: "巨", 2567: "巧", 2568: "左", 2569: "工", 2570: "巢", 2571: "巡", 2572: "州", 2573: "川", 2574: "巍", 2575: "嵌", 2576: "崭", 2577: "崩", 2578: "崖", 2579: "崔", 2580: "崎", 2581: "崇", 2582: "峻", 2583: "峰", 2584: "峭", 2585: "峦", 2586: "峡", 2587: "岸", 2588: "岳", 2589: "岭", 2590: "岩", 2591: "岛", 2592: "岗", 2593: "岖", 2594: "岔", 2595: "岂", 2596: "岁", 2597: "屿", 2598: "屹", 2599: "山", 2600: "屯", 2601: "履", 2602: "屡", 2603: "屠", 2604: "属", 2605: "展", 2606: "屑", 2607: "屏", 2608: "屎", 2609: "屋", 2610: "届", 2611: "屉", 2612: "屈", 2613: "居", 2614: "层", 2615: "屁", 2616: "局", 2617: "尿", 2618: "尾", 2619: "尽", 2620: "尼", 2621: "尺", 2622: "尸", 2623: "就", 2624: "尤", 2625: "尝", 2626: "尚", 2627: "尘", 2628: "尖", 2629: "尔", 2630: "少", 2631: "小", 2632: "尊", 2633: "尉", 2634: "将", 2635: "射", 2636: "封", 2637: "寿", 2638: "导", 2639: "寻", 2640: "寺", 2641: "对", 2642: "寸", 2643: "寨", 2644: "寥", 2645: "寡", 2646: "察", 2647: "寞", 2648: "寝", 2649: "寓", 2650: "寒", 2651: "富", 2652: "寇", 2653: "密", 2654: "寄", 2655: "寂", 2656: "宿", 2657: "宾", 2658: "宽", 2659: "容", 2660: "家", 2661: "宵", 2662: "宴", 2663: "害", 2664: "宰", 2665: "宫", 2666: "宪", 2667: "宦", 2668: "室", 2669: "宣", 2670: "客", 2671: "审", 2672: "宠", 2673: "实", 2674: "宝", 2675: "宜", 2676: "宛", 2677: "定", 2678: "宙", 2679: "官", 2680: "宗", 2681: "宏", 2682: "完", 2683: "宋", 2684: "安", 2685: "守", 2686: "宇", 2687: "宅", 2688: "它", 2689: "宁", 2690: "孽", 2691: "孵", 2692: "孩", 2693: "学", 2694: "孤", 2695: "季", 2696: "孟", 2697: "孝", 2698: "孙", 2699: "存", 2700: "字", 2701: "孕", 2702: "孔", 2703: "子", 2704: "嬉", 2705: "嫩", 2706: "嫡", 2707: "嫌", 2708: "嫉", 2709: "嫂", 2710: "嫁", 2711: "媳", 2712: "媚", 2713: "媒", 2714: "婿", 2715: "婶", 2716: "婴", 2717: "婚", 2718: "婉", 2719: "婆", 2720: "娶", 2721: "娱", 2722: "娩", 2723: "娜", 2724: "娘", 2725: "娇", 2726: "娄", 2727: "娃", 2728: "威", 2729: "姿", 2730: "姻", 2731: "姨", 2732: "姥", 2733: "姜", 2734: "姚", 2735: "委", 2736: "姓", 2737: "姑", 2738: "姐", 2739: "始", 2740: "姊", 2741: "姆", 2742: "妻", 2743: "妹", 2744: "妨", 2745: "妥", 2746: "妙", 2747: "妖", 2748: "妓", 2749: "妒", 2750: "妈", 2751: "妇", 2752: "妆", 2753: "妄", 2754: "如", 2755: "好", 2756: "她", 2757: "奸", 2758: "奶", 2759: "奴", 2760: "女", 2761: "奥", 2762: "奢", 2763: "奠", 2764: "套", 2765: "奖", 2766: "奕", 2767: "奔", 2768: "契", 2769: "奏", 2770: "奋", 2771: "奉", 2772: "奈", 2773: "奇", 2774: "奄", 2775: "夺", 2776: "夹", 2777: "夸", 2778: "夷", 2779: "头", 2780: "失", 2781: "夯", 2782: "央", 2783: "夭", 2784: "夫", 2785: "太", 2786: "天", 2787: "大", 2788: "够", 2789: "夜", 2790: "多", 2791: "外", 2792: "夕", 2793: "夏", 2794: "复", 2795: "备", 2796: "处", 2797: "壹", 2798: "壶", 2799: "壳", 2800: "声", 2801: "壮", 2802: "士", 2803: "壤", 2804: "壕", 2805: "壁", 2806: "墩", 2807: "墨", 2808: "增", 2809: "墙", 2810: "墓", 2811: "墅", 2812: "境", 2813: "填", 2814: "塞", 2815: "塘", 2816: "塔", 2817: "塑", 2818: "塌", 2819: "堵", 2820: "堰", 2821: "堪", 2822: "堤", 2823: "堡", 2824: "堕", 2825: "堆", 2826: "堂", 2827: "基", 2828: "培", 2829: "埠", 2830: "域", 2831: "城", 2832: "埋", 2833: "埃", 2834: "埂", 2835: "垮", 2836: "垫", 2837: "垦", 2838: "垢", 2839: "垛", 2840: "垒", 2841: "型", 2842: "垄", 2843: "垃", 2844: "垂", 2845: "坷", 2846: "坯", 2847: "坪", 2848: "坦", 2849: "坤", 2850: "坡", 2851: "坠", 2852: "坟", 2853: "坞", 2854: "坝", 2855: "坛", 2856: "坚", 2857: "块", 2858: "坑", 2859: "坐", 2860: "坏", 2861: "坎", 2862: "坊", 2863: "均", 2864: "址", 2865: "圾", 2866: "场", 2867: "地", 2868: "在", 2869: "圣", 2870: "土", 2871: "圈", 2872: "圆", 2873: "圃", 2874: "图", 2875: "国", 2876: "固", 2877: "围", 2878: "囱", 2879: "困", 2880: "园", 2881: "囤", 2882: "团", 2883: "因", 2884: "回", 2885: "四", 2886: "囚", 2887: "囊", 2888: "嚼", 2889: "嚷", 2890: "嚣", 2891: "嚎", 2892: "噪", 2893: "噩", 2894: "器", 2895: "嘿", 2896: "嘹", 2897: "嘶", 2898: "嘴", 2899: "嘲", 2900: "嘱", 2901: "嘉", 2902: "嘁", 2903: "嘀", 2904: "嗽", 2905: "嗦", 2906: "嗤", 2907: "嗡", 2908: "嗜", 2909: "嗓", 2910: "嗅", 2911: "喻", 2912: "喷", 2913: "喳", 2914: "喧", 2915: "喝", 2916: "喜", 2917: "喘", 2918: "喊", 2919: "喉", 2920: "喇", 2921: "善", 2922: "喂", 2923: "啼", 2924: "啸", 2925: "啰", 2926: "啦", 2927: "啥", 2928: "啤", 2929: "啡", 2930: "啊", 2931: "商", 2932: "啄", 2933: "啃", 2934: "唾", 2935: "唱", 2936: "唯", 2937: "售", 2938: "唬", 2939: "唧", 2940: "唤", 2941: "唠", 2942: "唐", 2943: "唉", 2944: "唇", 2945: "唆", 2946: "唁", 2947: "哼", 2948: "哺", 2949: "哲", 2950: "哮", 2951: "哭", 2952: "哪", 2953: "哩", 2954: "哨", 2955: "哥", 2956: "哟", 2957: "哗", 2958: "哑", 2959: "哎", 2960: "响", 2961: "哈", 2962: "哆", 2963: "哄", 2964: "品", 2965: "哀", 2966: "咽", 2967: "咸", 2968: "咳", 2969: "咱", 2970: "咬", 2971: "咪", 2972: "咨", 2973: "咧", 2974: "咙", 2975: "咖", 2976: "咕", 2977: "咒", 2978: "咐", 2979: "咏", 2980: "和", 2981: "咆", 2982: "命", 2983: "呼", 2984: "呻", 2985: "呵", 2986: "味", 2987: "周", 2988: "呢", 2989: "呜", 2990: "呛", 2991: "员", 2992: "呕", 2993: "呐", 2994: "告", 2995: "呈", 2996: "呆", 2997: "呀", 2998: "吼", 2999: "吻", 3000: "吹", 3001: "吸", 3002: "吵", 3003: "吴", 3004: "吱", 3005: "启", 3006: "吮", 3007: "吭", 3008: "听", 3009: "含", 3010: "吩", 3011: "吨", 3012: "吧", 3013: "否", 3014: "吠", 3015: "吟", 3016: "吞", 3017: "吝", 3018: "君", 3019: "吗", 3020: "吕", 3021: "吓", 3022: "向", 3023: "吐", 3024: "吏", 3025: "后", 3026: "名", 3027: "同", 3028: "吊", 3029: "吉", 3030: "合", 3031: "吆", 3032: "各", 3033: "吃", 3034: "吁", 3035: "叽", 3036: "叼", 3037: "叹", 3038: "司", 3039: "号", 3040: "叶", 3041: "右", 3042: "史", 3043: "台", 3044: "可", 3045: "叮", 3046: "叭", 3047: "召", 3048: "叫", 3049: "只", 3050: "叨", 3051: "另", 3052: "句", 3053: "古", 3054: "口", 3055: "叠", 3056: "叛", 3057: "叙", 3058: "变", 3059: "受", 3060: "取", 3061: "叔", 3062: "发", 3063: "反", 3064: "双", 3065: "友", 3066: "及", 3067: "叉", 3068: "又", 3069: "参", 3070: "叁", 3071: "县", 3072: "去", 3073: "厨", 3074: "厦", 3075: "厢", 3076: "原", 3077: "厚", 3078: "厘", 3079: "厕", 3080: "厌", 3081: "压", 3082: "厉", 3083: "历", 3084: "厅", 3085: "厂", 3086: "卿", 3087: "卸", 3088: "卷", 3089: "卵", 3090: "却", 3091: "即", 3092: "危", 3093: "印", 3094: "卫", 3095: "卧", 3096: "卦", 3097: "卤", 3098: "卢", 3099: "卡", 3100: "占", 3101: "卜", 3102: "博", 3103: "南", 3104: "卖", 3105: "单", 3106: "卓", 3107: "卒", 3108: "卑", 3109: "协", 3110: "华", 3111: "半", 3112: "午", 3113: "升", 3114: "千", 3115: "十", 3116: "匿", 3117: "匾", 3118: "医", 3119: "区", 3120: "匹", 3121: "匪", 3122: "匣", 3123: "匠", 3124: "匙", 3125: "北", 3126: "化", 3127: "匕", 3128: "匈", 3129: "匆", 3130: "包", 3131: "匀", 3132: "勿", 3133: "勾", 3134: "勺", 3135: "勤", 3136: "募", 3137: "勘", 3138: "勒", 3139: "勋", 3140: "勉", 3141: "勇", 3142: "勃", 3143: "势", 3144: "劳", 3145: "劲", 3146: "励", 3147: "劫", 3148: "努", 3149: "助", 3150: "动", 3151: "劣", 3152: "务", 3153: "加", 3154: "功", 3155: "办", 3156: "劝", 3157: "力", 3158: "劈", 3159: "剿", 3160: "割", 3161: "副", 3162: "剪", 3163: "剩", 3164: "剧", 3165: "剥", 3166: "剖", 3167: "剔", 3168: "剑", 3169: "前", 3170: "削", 3171: "剃", 3172: "剂", 3173: "刽", 3174: "刻", 3175: "刺", 3176: "刹", 3177: "券", 3178: "刷", 3179: "制", 3180: "到", 3181: "刮", 3182: "别", 3183: "利", 3184: "刨", 3185: "判", 3186: "删", 3187: "初", 3188: "创", 3189: "刚", 3190: "则", 3191: "刘", 3192: "列", 3193: "划", 3194: "刑", 3195: "刊", 3196: "切", 3197: "分", 3198: "刃", 3199: "刁", 3200: "刀", 3201: "凿", 3202: "函", 3203: "击", 3204: "出", 3205: "凹", 3206: "凸", 3207: "凶", 3208: "凳", 3209: "凰", 3210: "凯", 3211: "凭", 3212: "凫", 3213: "凤", 3214: "凡", 3215: "几", 3216: "凝", 3217: "凛", 3218: "凑", 3219: "减", 3220: "凌", 3221: "凉", 3222: "准", 3223: "凄", 3224: "净", 3225: "冻", 3226: "冷", 3227: "冶", 3228: "况", 3229: "决", 3230: "冲", 3231: "冰", 3232: "冯", 3233: "冬", 3234: "冤", 3235: "冠", 3236: "农", 3237: "军", 3238: "写", 3239: "冗", 3240: "冕", 3241: "冒", 3242: "再", 3243: "册", 3244: "冈", 3245: "内", 3246: "冀", 3247: "兽", 3248: "兼", 3249: "养", 3250: "典", 3251: "具", 3252: "其", 3253: "兵", 3254: "兴", 3255: "关", 3256: "共", 3257: "兰", 3258: "六", 3259: "公", 3260: "八", 3261: "全", 3262: "入", 3263: "兢", 3264: "兜", 3265: "党", 3266: "兔", 3267: "兑", 3268: "免", 3269: "克", 3270: "光", 3271: "先", 3272: "兆", 3273: "充", 3274: "兄", 3275: "元", 3276: "允", 3277: "儿", 3278: "儡", 3279: "儒", 3280: "僻", 3281: "僵", 3282: "僧", 3283: "僚", 3284: "像", 3285: "傻", 3286: "傲", 3287: "催", 3288: "储", 3289: "傍", 3290: "傅", 3291: "傀", 3292: "偿", 3293: "偷", 3294: "偶", 3295: "健", 3296: "停", 3297: "做", 3298: "偏", 3299: "偎", 3300: "假", 3301: "倾", 3302: "值", 3303: "债", 3304: "倦", 3305: "倡", 3306: "借", 3307: "倚", 3308: "候", 3309: "倘", 3310: "倔", 3311: "倒", 3312: "倍", 3313: "俺", 3314: "俱", 3315: "俯", 3316: "修", 3317: "俭", 3318: "俩", 3319: "信", 3320: "保", 3321: "俘", 3322: "俗", 3323: "俐", 3324: "俏", 3325: "俊", 3326: "俄", 3327: "促", 3328: "便", 3329: "侵", 3330: "侯", 3331: "侮", 3332: "侨", 3333: "侧", 3334: "侦", 3335: "侥", 3336: "侣", 3337: "侠", 3338: "依", 3339: "供", 3340: "侍", 3341: "例", 3342: "侈", 3343: "侄", 3344: "使", 3345: "佳", 3346: "佩", 3347: "佣", 3348: "你", 3349: "作", 3350: "佛", 3351: "余", 3352: "何", 3353: "体", 3354: "佑", 3355: "住", 3356: "低", 3357: "位", 3358: "但", 3359: "佃", 3360: "似", 3361: "伺", 3362: "伸", 3363: "伶", 3364: "伴", 3365: "估", 3366: "伯", 3367: "伪", 3368: "伦", 3369: "伤", 3370: "传", 3371: "伟", 3372: "伞", 3373: "会", 3374: "伙", 3375: "优", 3376: "众", 3377: "休", 3378: "伐", 3379: "伏", 3380: "伍", 3381: "伊", 3382: "企", 3383: "仿", 3384: "份", 3385: "任", 3386: "价", 3387: "件", 3388: "仲", 3389: "仰", 3390: "们", 3391: "仪", 3392: "以", 3393: "令", 3394: "代", 3395: "仙", 3396: "付", 3397: "仗", 3398: "他", 3399: "仔", 3400: "仓", 3401: "仑", 3402: "从", 3403: "仍", 3404: "介", 3405: "今", 3406: "仇", 3407: "仆", 3408: "仅", 3409: "仁", 3410: "什", 3411: "亿", 3412: "人", 3413: "亲", 3414: "亮", 3415: "亭", 3416: "京", 3417: "享", 3418: "亩", 3419: "产", 3420: "亦", 3421: "亥", 3422: "交", 3423: "亡", 3424: "些", 3425: "亚", 3426: "井", 3427: "五", 3428: "互", 3429: "云", 3430: "亏", 3431: "于", 3432: "二", 3433: "事", 3434: "争", 3435: "予", 3436: "了", 3437: "乾", 3438: "乳", 3439: "乱", 3440: "买", 3441: "书", 3442: "乡", 3443: "习", 3444: "也", 3445: "乞", 3446: "九", 3447: "乙", 3448: "乘", 3449: "乖", 3450: "乔", 3451: "乓", 3452: "乒", 3453: "乐", 3454: "乏", 3455: "乎", 3456: "乍", 3457: "乌", 3458: "之", 3459: "义", 3460: "么", 3461: "久", 3462: "乃", 3463: "举", 3464: "丽", 3465: "主", 3466: "为", 3467: "丹", 3468: "丸", 3469: "临", 3470: "串", 3471: "丰", 3472: "中", 3473: "个", 3474: "丧", 3475: "严", 3476: "两", 3477: "丢", 3478: "丝", 3479: "东", 3480: "丛", 3481: "业", 3482: "丙", 3483: "丘", 3484: "世", 3485: "且", 3486: "专", 3487: "丑", 3488: "丐", 3489: "与", 3490: "不", 3491: "下", 3492: "上", 3493: "三", 3494: "丈", 3495: "万", 3496: "七", 3497: "丁"}
	rand.Seed(time.Now().UnixNano())
	str := ""
	for i := 0; i < leng; i++ {
		k := rand.Intn(3497)
		str = str + zhmap[k]
	}
	return str
}
