# git | badges

Generate custom badges: [shields.io](https://shields.io/)

## Jest Coverage

Script developed in shell and node to cat the coverage data and stringify. After prepare the content execute curl to Gist.

```bash
content=`node -pe 'JSON.stringify(process.argv[1])' "$(cat coverage-summary.json)"`
curl -X PATCH  -H "Accept: application/vnd.github.v3+json" -u leoduprates:${{ secrets.GIST_SECRET }} https://api.github.com/gists/80e2092dfa47c651009be8e7cbb3ef84 -d '{"files":{"javascript-testing-coverage.json":{"content": '"$content"'}}}'
```

## Jest Test Report

Script developed in shell and node to cat the report data and replace the value in success field and stringify.

```bash
content=`node -pe 'const content = JSON.parse(process.argv[1]); content.success = content.success == true ? "passing" : "failing"; JSON.stringify(content) ' "$(cat tests/reports/backend_report.json)"`
```

## Badge Colors

### Github Style

```text
labelColor (left): 3e464f
success color (rigth): 31c854
```

### Markdown Badges

- <https://github.com/Ileriayo/markdown-badges>
- <https://badgen.net/>