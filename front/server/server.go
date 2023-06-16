package server

import (
)

func Init() {
    r := Routes()
    r.Run(":9090")
}
