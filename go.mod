module NTNU_Prj/WebSocket

go 1.12

require github.com/gorilla/websocket v1.4.1

require (
	NTNU_Prj/DeCodeURInterface v0.0.0
	NTNU_Prj/DualArmControl v0.0.0
	NTNU_Prj/UR3Handler v0.0.0
)

replace (
	NTNU_Prj/DeCodeURInterface => ../DeCodeURInterface
	NTNU_Prj/DualArmControl => ../DualArmControl
	NTNU_Prj/UR3Handler => ../UR3Handler
)
