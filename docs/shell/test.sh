
get_all_user(){
  cat /etc/passwd | cut -f 1 -d : &&  return
}

get_opsy() {
    [ -f /etc/redhat-release ] && awk '{print $0}' /etc/redhat-release && return
    [ -f /etc/os-release ] && awk -F'[= "]' '/PRETTY_NAME/{print $3,$4,$5}' /etc/os-release && return
    [ -f /etc/lsb-release ] && awk -F'[="]+' '/DESCRIPTION/{print $2}' /etc/lsb-release && return
}

get_plan_task(){
  users=$( get_all_user )
  declare -A plan_task
  for user in $users; do
    res=$(crontab -l -u $user)
    plan_task[user]= $res
  done
  echo $plan_task
  return ${plan_task[*]}
}

opsy=$( get_opsy )
plan=$( get_plan_task )

echo "{ \"os\": \""$opsy"\", \"plan\":\""$plan"\" }"
