package main

import (
        "gopkg.in/qml.v1"
        "log"
)

func main() {
        err := qml.Run(run)
        if (err != nil) {
                log.Fatal(err)
        }
}

func run() error {
        engine := qml.NewEngine()
        component, err := engine.LoadFile("share/net-info/main.qml")
        if err != nil {
                return err
        }

        ctrl := Control{Message: "Active interfaces"}
        context := engine.Context()
        context.SetVar("ctrl", &ctrl)
        ctrl.GetInterfaces()
        win := component.CreateWindow(nil)
        ctrl.Root = win.Root()
        win.Show()
        win.Wait()
        return nil
}

type Control struct {
	Root    qml.Object
  Message string
  Output string
}

func (ctrl *Control) GetInterfaces() {
  log.Print("fetch interfaces...")
  go func() {
    ctxt := PhoneInterfacesString()
    ctrl.Output = ctxt
    qml.Changed(ctrl, &ctrl.Output)
  }()
}

