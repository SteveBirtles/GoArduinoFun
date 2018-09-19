package main

import (
        "time"
        "gobot.io/x/gobot"
        "gobot.io/x/gobot/drivers/gpio"
        "gobot.io/x/gobot/platforms/firmata"
)

func main() {

        firmataAdaptor := firmata.NewAdaptor("/dev/ttyACM0")

        led1 := gpio.NewDirectPinDriver(firmataAdaptor, "7")
        led2 := gpio.NewDirectPinDriver(firmataAdaptor, "4")
        led3 := gpio.NewDirectPinDriver(firmataAdaptor, "3")
        led4 := gpio.NewDirectPinDriver(firmataAdaptor, "8")
        pwm := gpio.NewDirectPinDriver(firmataAdaptor, "11")
        buzz := gpio.NewDirectPinDriver(firmataAdaptor, "12")

        r, g, b, y, z := 1, 3, 5, 7, 8

        work := func() {

          led1.On()
          led2.On()
          led3.On()
          led4.On()

          gobot.Every(80*time.Millisecond, func() {

            if z--; z <= 0 {
              z += 8
              buzz.On()
            } else if z == 7 {
              buzz.Off()
            } else if z == 5 {
              pwm.PwmWrite(128)
            } else if z == 4 {
              pwm.Off()
            }

            if r--; r <= 0 {
              r += 8
              if val, _ := led1.DigitalRead(); val == 1 {
                led1.Off()
              } else {
                led1.On()
              }
            }

            if g--; g <= 0 {
              g += 8
              if val, _ := led2.DigitalRead(); val == 1 {
                led2.Off()
              } else {
                led2.On()
              }
            }

            if b--; b <= 0 {
              b += 8
              if val, _ := led3.DigitalRead(); val == 1 {
                led3.Off()
              } else {
                led3.On()
              }
            }

            if y--; y <= 0 {
              y += 8
              if val, _ := led4.DigitalRead(); val == 1 {
                led4.Off()
              } else {
                led4.On()
              }
            }

          })

        }

        robot := gobot.NewRobot("bot",
                []gobot.Connection{firmataAdaptor},
                []gobot.Device{led1},
                []gobot.Device{led2},
                []gobot.Device{led3},
                []gobot.Device{led4},
                []gobot.Device{buzz},
                work,
        )

        robot.Start()
}
