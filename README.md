# random_arXiv
arXivの論文をランダムで表示するコマンドラインツール 

## デモ(DEMO)
```
$ random_arXiv view
```

```
Composition of Saliency Metrics for Channel Pruning with a Myopic Oracle

2020-04-03T11:29:41Z
2020-04-03T11:29:41Z

  The computation and memory needed for Convolutional Neural Network (CNN)
inference can be reduced by pruning weights from the trained network. Pruning
is guided by a pruning saliency, which heuristically approximates the change in
the loss function associated with the removal of specific weights. Many pruning
signals have been proposed, but the performance of each heuristic depends on
the particular trained network. This leaves the data scientist with a difficult
choice. When using any one saliency metric for the entire pruning process, we
run the risk of the metric assumptions being invalidated, leading to poor
decisions being made by the metric. Ideally we could combine the best aspects
of different saliency metrics. However, despite an extensive literature review,
we are unable to find any prior work on composing different saliency metrics.
The chief difficulty lies in combining the numerical output of different
saliency metrics, which are not directly comparable.
  We propose a method to compose several primitive pruning saliencies, to
exploit the cases where each saliency measure does well. Our experiments show
that the composition of saliencies avoids many poor pruning choices identified
by individual saliencies. In most cases our method finds better selections than
even the best individual pruning saliency.

David Gregg

http://arxiv.org/abs/2004.03376v1
http://arxiv.org/pdf/2004.03376v1

cs.CV
cs.CV
cs.LG
stat.ML
```

## 機能(Features)
* AI関連のカテゴリの論文を表示
* 乱数を用いて論文のインデックスを指定
* タイトルや著者，要約などを色を変えてターミナルに出力

## 使い方(Usage)

## 環境(Requirement, Environment)
* Go
* cobra

## インストール(Installation)

## 注意事項(Note)

## 文責(Author)
* mkdk09
* mkdk099@gmail.com

## ライセンス
This code is under [MIT license](https://en.wikipedia.org/wiki/MIT_License).

## 参考文献(References)
### cobra
https://taisablog.com/archives/1908  
https://blog.engineer.adways.net/entry/advent_calendar_2018/18  
https://uzimihsr.github.io/post/2020-09-03-golang-cobra/  
### XML
https://qiita.com/ytkhs/items/948f516ec82c82eaa882  
https://www.onlinetool.io/xmltogo/  
### API
https://qiita.com/KMD/items/bd59f2db778dd4bf6ed2  
https://webbigdata.jp/ai/post-4764  
### Terminal
https://github.com/fatih/color  
https://github.com/logrusorgru/aurora  
### Test
https://qiita.com/tkit/items/3cdeafcde2bd98612428  
https://stackoverflow.com/questions/35827147/cobra-viper-golang-how-to-test-subcommands  
