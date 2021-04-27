package snakesnow

var ServerId = 1

var UserIdWorkMap = make(map[int]*IdWorker)
var GroupIdWorkMap = make(map[int]*IdWorker)
var DeviceIdWorkMap = make(map[int]*IdWorker)

func RandUserId() (int64, error) {
	value, ok := UserIdWorkMap[ServerId]
	if ok {
		nid, err := value.NextId()
		return nid, err
	} else {
		iw, err := NewIdWorker(int64(ServerId))
		if err != nil {
			return 0, err
		}
		nid, err := iw.NextId()
		UserIdWorkMap[ServerId] = iw
		return nid, err
	}
}
func RandGroupId() (int64, error) {
	value, ok := UserIdWorkMap[ServerId]
	if ok {
		nid, err := value.NextId()
		return nid, err
	} else {
		iw, err := NewIdWorker(int64(ServerId))
		if err != nil {
			return 0, err
		}
		nid, err := iw.NextId()
		UserIdWorkMap[ServerId] = iw
		return nid, err
	}
}
func RandDeviceId() (int64, error) {
	value, ok := UserIdWorkMap[ServerId]
	if ok {
		nid, err := value.NextId()
		return nid, err
	} else {
		iw, err := NewIdWorker(int64(ServerId))
		if err != nil {
			return 0, err
		}
		nid, err := iw.NextId()
		UserIdWorkMap[ServerId] = iw
		return nid, err
	}
}
