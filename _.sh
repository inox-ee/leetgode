leetgode () {
        if [ $# -ne 2 ]
        then
                echo "指定された引数は $# 個です。" >&2
                echo "実行するには問題番号とタイトルを指定してください。" >&2
                return
        fi
        CURRENTPATH=$(pwd) 
        if [ ${CURRENTPATH##*/} != leetgode ]
        then
                echo "現在のディレクトリは ${CURRENTPATH} です。" >&2
                echo "実行するには path/to/leetgode に移動してください。" >&2
                return
        fi
        {
                mkdir -p -p "./leetcode/$1.$2"
                cd $_
                touch "$2.go"
                curl -LO "https://raw.githubusercontent.com/aQuaYi/LeetCode-in-Go/master/Algorithms/$1.$2/$2_test.go"
                echo "create $1.$2/$2.go, $2_test.go"
        } always {
                if catch '*'
                then
                        case $CAUGHT in
                                (*) echo "error" ;;
                        esac
                fi
        }
        return
}
