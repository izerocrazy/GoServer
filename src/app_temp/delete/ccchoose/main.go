package main

import "fmt"
import "math/rand"

var OLD = 1
var NEW = 2

var PERSONAL = 1
var MACHINE = 2

var MAINModen int

var MAXCOUNT = 10000

func main() {
    fmt.Printf("请选择：\n")
    fmt.Printf("1：人工输入\n")
    fmt.Printf("2：机器输入\n")

    fmt.Scanf("%d", &MAINModen)

    nCount := 0
    nOLDWin := 0
    nNEWWin := 0

    for nCount < MAXCOUNT {
        nChoose, bWin := OneChoose(nCount)
        if nChoose == OLD && bWin == true {
            nOLDWin = nOLDWin + 1
        }

        if nChoose == NEW && bWin == true {
            nNEWWin = nNEWWin + 1
        }

        nCount = nCount + 1

        fmt.Printf("当前一共跑了 %d 次 \n", nCount)
        fmt.Printf(">>>>>>> 当前坚持并且胜利的概率为 %f \n", float32(nOLDWin) / float32(nCount))
        fmt.Printf(">>>>>>> 不坚持并胜利的概率为 %f \n\n", float32(nNEWWin) / float32(nCount))
    }
}

func OneChoose(nSeed int) (int, bool){
    rand.Seed(int64(nSeed))
    answers := []string{"A", "B", "C"}

    nRigheIndex := rand.Intn(len(answers))

    var nDeleteIndex int
    for{
        nDeleteIndex = rand.Intn(len(answers))
        if nDeleteIndex != nRigheIndex {
            break
        }
    }

    var nFirstChoose int
    for {
        nFirstChoose = rand.Intn(len(answers))
        if nFirstChoose != nDeleteIndex {
            break
        }
    }

    fmt.Printf("有选项 A B C，你预先选择了 %s\n", answers[nFirstChoose])

    var nFinalChoose int
    var nChoose int
    for {
        fmt.Printf("现在群猪帮你去掉了一个错误答案 %s，请问你还坚持原本选择的答案 %s 吗？yes 输入 1，no 输入 2:\n", answers[nDeleteIndex], answers[nFirstChoose])

        if MAINModen == PERSONAL {
            fmt.Scanf("%d", &nChoose)
        } else {
            rand.Seed(int64(nSeed * nSeed))
            var nBig int
            if NEW > OLD {
                nBig = NEW
            } else {
                nBig = OLD
            }
            nChoose = rand.Intn(nBig) + 1

            fmt.Printf("MACHINE choose %d\n", nChoose)
        }

        if nChoose == OLD {
            nFinalChoose = nFirstChoose
            break
        } else if(nChoose == NEW) {
            for i := 0; i < 3; i++ {
                if i != nFirstChoose && i != nDeleteIndex {
                    nFinalChoose = i
                    break
                }
            }
            break
        } else {
            fmt.Println("err...... Let's try again\n")
        }
    }

    var bWin bool
    if nFinalChoose == nRigheIndex {
        bWin = true
    } else {
        bWin = false
    }

    //fmt.Print(nFirstChoose, nDeleteIndex, nRigheIndex, nFinalChoose)
    fmt.Printf("正确答案是 %s，你最终选择是 %s \n", answers[nRigheIndex], answers[nFinalChoose])

    fmt.Printf("\n")

    return nChoose, bWin
}
