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
                local response=$(curl -LOf -w '%{http_code}\n' "https://raw.githubusercontent.com/aQuaYi/LeetCode-in-Go/master/Algorithms/$1.$2/$2_test.go")
                if [ $response -eq 200 ]; then
                        mkdir -p "./leetcode/$1.$2"
                        mv ./$2_test.go $_/
                        cd $_
                        touch "$2.go" && touch "README.md" && echo "create $1.$2/$2.go, $2_test.go"
                else
                        throw "CurlError"
                fi
        } always {
                if catch '*'; then
                case $CAUGHT in
                        (*)
                        echo $CAUGHT
                        ;;
                esac
                fi
        }
        return
}
