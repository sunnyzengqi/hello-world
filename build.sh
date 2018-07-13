#!/bin/bash
# go默认使用go1.4
# 如果需要使用go1.6 ，请打开下面两行注释：

#export GOROOT=/usr/local/go
#服务器上
export GOROOT=/usr/lib/golang
export PATH=$GOROOT/bin:$PATH

workspace=$(cd $(dirname $0) && pwd -P)

tempbuildpath="$workspace/temp_zaixianshang_build_dir"
tempcodepath="$tempbuildpath/src/hello-world"

mkdir -p $tempcodepath
mkdir $tempbuildpath/bin


export GOPATH=$tempbuildpath
export PATH=${GOPATH}/bin:$PATH

curl https://git.xiaojukeji.com/lego/tools/raw/master/glide/get | sh

mkdir -p $tempcodepath
mkdir $tempcodepath/bin

cp -rf `ls ./ | grep -E -v "^(temp_zaixianshang_build_dir)$"` $tempcodepath/

cd $tempcodepath
module=nova-toc-backend2
control=control.sh
output="output"

#git config --global url."git@git.xiaojukeji.com:".insteadOf "https://git.xiaojukeji.com/"

echo $GOPATH


git config --global url."git@git.xiaojukeji.com:".insteadOf "https://git.xiaojukeji.com/"
#chmod +x glide.sh
#sh glide.sh

#glide init
pwd
glide install --vendor-override

go build -o $module DingTicker.go    #编译目标文件,需要编译那个主类就写哪个
ret=$?
if [ $ret -ne 0 ];then
    echo "===== $module build failure ====="
    exit $ret
else
    echo -n "===== $module build successfully! ====="
fi

cp ./${module} $workspace/
#rm -rf $workspace/vendor
mv ./vendor $workspace/
cd $workspace

rm -rf $tempbuildpath
rm -rf $output
mkdir -p $output
mkdir -p $output/bin
mkdir -p $output/conf
mkdir -p $output/logs

# 填充output目录, output的内容即为待部署内容
    (
        cp -rf ./${control} ${output}/ && chmod +x ${output}/${control} &&    # 拷贝部署脚本control.sh至output目录
        mv ${module} ${output}/bin/ &&        # 移动需要部署的文件到output目录下
	#cp -rf ./cfg ${output}/cfg/ &&
        echo -e "===== Generate output ok ====="
    ) || { echo -e "===== Generate output failure ====="; exit 2; } # 填充output目录失败后, 退出码为 非0

