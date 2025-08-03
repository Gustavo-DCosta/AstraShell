commit_message = Read-Host "What is your commit message?: "

git add .
git commit -m "$commit_message"
git push origin main