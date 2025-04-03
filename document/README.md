### Git 命令
```sh
git branch # 查看当前分支
git checkout main # 切换到 main 分支
git branch -m master main      # 本地分支重命名
git push -u origin main        # 推送新分支
git remote show origin        # 查看关联关系
git branch -vv                # 检查跟踪状态

# 将本地 master 分支推送到远程 main 分支
git push origin master:main

# 强制覆盖远程 main 分支
git push -f origin master:main

git push origin --delete master  # 删除远程 master 分支

# 删除远程旧分支
git push origin --delete feature/old

# 推送新分支到远程
git push -u origin feature/new
```