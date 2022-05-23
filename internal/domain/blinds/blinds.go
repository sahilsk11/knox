package blinds

type BlindsName string

const (
	Blinds_BigWindowFirstThird   BlindsName = "BLINDS_BIG_WINDOW_FIRST_THIRD"
	Blinds_BigWindowSecondThirds BlindsName = "BIG_WINDOW_SECOND_THIRDS"
	Blinds_Door                  BlindsName = "BLINDS_DOOR"
	Blinds_LivingRoom            BlindsName = "BLINDS_LIVING_ROOM"
)

type MotionType string

const (
	MotionType_Sequential MotionType = "MOTIONTYPE_SEQUENTIAL"
	MotionType_Concurrent MotionType = "MOTIONTYPE_CONCURRENT"
)
