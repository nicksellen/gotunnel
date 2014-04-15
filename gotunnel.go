package main

import (
  "log"
  "os/exec"
  "fmt"
  "io/ioutil"
  "os"
  "os/signal"
  "launchpad.net/goyaml"
  "text/tabwriter"
  "github.com/wsxiaoys/terminal/color"
)

type Tunnel struct {
  Local int `yaml:"local"`
  Remote int `yaml:"remote"`
  Protocol string `yaml:"protocol"`
}

type Conf struct {
  Hosts map[string]map[string]Tunnel `yaml:"hosts"`
}

func main() {

  bytes, err := ioutil.ReadFile(os.Args[1])

  if err != nil {
    log.Fatal(err)
  }

  var conf Conf
  goyaml.Unmarshal(bytes, &conf);

  w := new(tabwriter.Writer)
  w.Init(os.Stdout, 20, 4, 2, ' ', 0)

  color.Printf("@{!}%-14s %-26s %-10s %-10s\n", "host", "service", "local", "remote")

  for host, tunnels := range conf.Hosts {
    args := []string{}
    for name, Tunnel := range tunnels {

      args = append(args, "-L", fmt.Sprintf("%d:localhost:%d", Tunnel.Local, Tunnel.Remote))

      fmt.Printf("%-14s %-26s %-10d %-10d", host, name, Tunnel.Local, Tunnel.Remote)

      if Tunnel.Protocol == "http" {
        fmt.Printf(" http://localhost:%d", Tunnel.Local)
      }

      fmt.Printf("\n")

    }
    args = append(args, "-N", host)
    w.Flush()
    go func() {
      cmd := exec.Command("ssh", args...)    
      err = cmd.Start()
      if err != nil {
        log.Fatal(err)
      }
    }()
  }

  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt, os.Kill)
  <-c
  fmt.Println("exiting")
}
