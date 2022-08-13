
get_all_user(){
  [ -f /etc/passwd ] && cat /etc/passwd | cut -f 1 -d : &&  return
}

get_opsy() {
    [ -f /etc/redhat-release ] && awk '{print $0}' /etc/redhat-release && return
    [ -f /etc/os-release ] && awk -F'[= "]' '/PRETTY_NAME/{print $3,$4,$5}' /etc/os-release && return
    [ -f /etc/lsb-release ] && awk -F'[="]+' '/DESCRIPTION/{print $2}' /etc/lsb-release && return
}
declare -A plan_task
get_plan_task(){
  users=$( get_all_user )

  for user in $users; do
    `crontab -l -u $user >/dev/null 2>&1`
    if [[ $? -eq 0 ]];then
      plan_task[$user]=`crontab -l -u $user`
    fi
  done
}

opsy=$( get_opsy )
get_plan_task

#echo "{ \"os\": \""$opsy"\", \"plan\":\""${plan_task[*]}"\" }"
echo "{ \"key\": \"os\", \"name\": \"操作系统\", \"value\": \""$opsy"\" }||{ \"key\": \"plan\", \"name\": \"计划任务\", \"value\": \""${plan_task[*]}"\" }"
