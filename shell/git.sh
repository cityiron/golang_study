echo "执行的文件名：$0";

run_flag=true

if [ -n "$1" ]; then
  echo "包含第一个参数: $1"
else
  run_flag=false
  echo "没有包含第一参数"
fi

if [ "$run_flag" = true ]; then
  eval "$(ssh-agent -s)"

  ssh-add $1

  ssh -T git@github.com
fi


