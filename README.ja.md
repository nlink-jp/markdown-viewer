# markdown-viewer

セキュアなシングルバイナリMarkdownビューア＆ファイルブラウザ。GFM、Mermaidダイアグラム、シンタックスハイライト対応。

## 機能

- **2ペインUI**: 左に展開可能なファイルツリー、右にコンテンツビューア
- **シングルバイナリ**: UI資材（CSS, JS, HTMLテンプレート）はすべて埋め込み、外部依存なし
- **Markdownレンダリング**: GitHub風スタイル、シンタックスハイライト、Mermaidダイアグラム対応
- **デフォルトでセキュア**: ディレクトリトラバーサル防止、XSS対策のHTMLサニタイズ
- **グレースフルシャットダウン**: UIボタンまたは `Ctrl+C`

## インストール

[リリースページ](https://github.com/nlink-jp/markdown-viewer/releases)からプラットフォーム別の `mdv` バイナリをダウンロード。

```sh
unzip mdv-<version>-<os>-<arch>.zip
mv mdv /usr/local/bin/
```

## 設定

`mdv` は `config.json` ファイルまたはコマンドラインフラグで設定可能。優先順位（後のものが優先）：

1. デフォルト値
2. `$HOME/.config/mdv/config.json`
3. カレントディレクトリの `config.json`
4. 環境変数（`MDV_` プレフィックス、例: `MDV_PORT=9000`）
5. コマンドラインフラグ

`config.json.example` を参照。

## 使い方

Markdownファイルを含むディレクトリで `mdv` を実行：

```sh
mdv [flags]
```

ブラウザで `http://127.0.0.1:8080` を開く。

### フラグ

| フラグ | デフォルト | 説明 |
|--------|-----------|------|
| `-p, --port` | `8080` | 待受ポート |
| `-o, --open` | `false` | 起動時にブラウザを自動オープン |
| `-d, --dir` | `.` | 配信ルートディレクトリ |
| `--version` | | バージョン表示 |

## ビルド

Go 1.24 以降が必要。

```sh
git clone https://github.com/nlink-jp/markdown-viewer.git
cd markdown-viewer
make build        # ビルド → dist/mdv
make build-all    # クロスコンパイル → dist/mdv-<os>-<arch>
make test         # テスト実行
make clean        # dist/ 削除
```

## 参照

- [CHANGELOG.md](CHANGELOG.md)
- [NOTICE.md](NOTICE.md) — サードパーティライセンス表記
