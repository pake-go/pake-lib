package config

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	cfg := New()

	expectedCurrent := make(map[string]string)
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}
	expectedOld := make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedSetTempAgeTracker := make(map[string]int)
	if !reflect.DeepEqual(cfg.setTemporarilyAgeTracker, expectedSetTempAgeTracker) {
		t.Errorf("Expected %+q but got %+q",
			expectedSetTempAgeTracker,
			cfg.setTemporarilyAgeTracker)
	}
	expectedSetTemporarilyAge := 1
	if cfg.setTemporarilyAge != expectedSetTemporarilyAge {
		t.Errorf("Expected %d but got %d",
			expectedSetTemporarilyAge,
			cfg.setTemporarilyAge)
	}
}

func TestWithSetTemporarily(t *testing.T) {
	cfg := WithSetTemporarilyAge(3)

	expectedCurrent := make(map[string]string)
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}
	expectedOld := make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedSetTempAgeTracker := make(map[string]int)
	if !reflect.DeepEqual(cfg.setTemporarilyAgeTracker, expectedSetTempAgeTracker) {
		t.Errorf("Expected %+q but got %+q",
			expectedSetTempAgeTracker,
			cfg.setTemporarilyAgeTracker)
	}
	expectedSetTemporarilyAge := 3
	if cfg.setTemporarilyAge != expectedSetTemporarilyAge {
		t.Errorf("Expected %d but got %d", expectedSetTemporarilyAge, cfg.setTemporarilyAge)
	}
}

func TestGet_keyexists(t *testing.T) {
	cfg := New()
	cfg.current["key"] = "value"

	value, err := cfg.Get("key")
	if err != nil {
		t.Error("Was not able to retrieve the value for `key`")
	}
	if value != "value" {
		t.Errorf("Expected value but got %s", value)
	}
}

func TestGet_nonexistentkey(t *testing.T) {
	cfg := New()

	if _, err := cfg.Get("nonExistentKey"); err == nil {
		t.Error("Should not able to retrieve any value")
	}
}

func TestSetTemporarily_keyexists(t *testing.T) {
	cfg := New()

	cfg.SetTemporarily("key", "value")
	expectedOld := make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedCurrent := make(map[string]string)
	expectedCurrent["key"] = "value"
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}

}

func TestSetTemporarily_nonexistentkey(t *testing.T) {
	cfg := New()
	cfg.current["key"] = "value"

	cfg.SetTemporarily("key", "value2")
	expectedOld := make(map[string]string)
	expectedOld["key"] = "value"
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedCurrent := make(map[string]string)
	expectedCurrent["key"] = "value2"
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}
}

func TestSetPermanently_keyexists(t *testing.T) {
	cfg := New()

	cfg.SetPermanently("key", "value")
	expectedOld := make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedCurrent := make(map[string]string)
	expectedCurrent["key"] = "value"
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}
}

func TestSetPermanently_nonexistentkey(t *testing.T) {
	cfg := New()
	cfg.current["key"] = "value"

	cfg.SetPermanently("key", "value2")
	expectedOld := make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedCurrent := make(map[string]string)
	expectedCurrent["key"] = "value2"
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}
}

func TestReset_emptyconfig(t *testing.T) {
	cfg := New()
	cfg.Reset()

	expectedOld := make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedCurrent := make(map[string]string)
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}
}

func TestReset_settemporarilydone(t *testing.T) {
	cfg := New()
	cfg.SetTemporarily("key", "value")
	cfg.Reset()

	expectedOld := make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedCurrent := make(map[string]string)
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}
}

func TestReset_setpermanentlydone(t *testing.T) {
	cfg := New()
	cfg.SetPermanently("key", "value")
	cfg.Reset()

	expectedOld := make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedCurrent := make(map[string]string)
	expectedCurrent["key"] = "value"
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}
}

func TestSmartReset_emptyconfigdefaultage(t *testing.T) {
	cfg := New()

	cfg.SmartReset()
	expectedOld := make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedCurrent := make(map[string]string)
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}

	cfg.SmartReset()
	expectedOld = make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedCurrent = make(map[string]string)
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}
}

func TestSmartReset_emptyconfigmodifiedage(t *testing.T) {
	cfg := WithSetTemporarilyAge(2)

	cfg.SmartReset()
	expectedOld := make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedCurrent := make(map[string]string)
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}

	cfg.SmartReset()
	expectedOld = make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedCurrent = make(map[string]string)
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}

	cfg.SmartReset()
	expectedOld = make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedCurrent = make(map[string]string)
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}
}

func TestSmartReset_settemporarilydonedefaultage(t *testing.T) {
	cfg := New()
	cfg.SetTemporarily("key", "value")

	cfg.SmartReset()
	expectedOld := make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedCurrent := make(map[string]string)
	expectedCurrent["key"] = "value"
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}

	cfg.SmartReset()
	expectedOld = make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedCurrent = make(map[string]string)
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}
}

func TestSmartReset_settemporarilydonemodifiedage(t *testing.T) {
	cfg := WithSetTemporarilyAge(2)
	cfg.SetTemporarily("key", "value")

	cfg.SmartReset()
	expectedOld := make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedCurrent := make(map[string]string)
	expectedCurrent["key"] = "value"
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}

	cfg.SmartReset()
	expectedOld = make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedCurrent = make(map[string]string)
	expectedCurrent["key"] = "value"
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}

	cfg.SmartReset()
	expectedOld = make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedCurrent = make(map[string]string)
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}
}

func TestSmartReset_setpermanentlydonedefaultage(t *testing.T) {
	cfg := New()
	cfg.SetPermanently("key", "value")

	cfg.SmartReset()
	expectedOld := make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedCurrent := make(map[string]string)
	expectedCurrent["key"] = "value"
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}

	cfg.SmartReset()
	expectedOld = make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedCurrent = make(map[string]string)
	expectedCurrent["key"] = "value"
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}

}

func TestSmartReset_setpermanentlydonemodifiedage(t *testing.T) {
	cfg := WithSetTemporarilyAge(2)
	cfg.SetPermanently("key", "value")

	cfg.SmartReset()
	expectedOld := make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedCurrent := make(map[string]string)
	expectedCurrent["key"] = "value"
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}

	cfg.SmartReset()
	expectedOld = make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedCurrent = make(map[string]string)
	expectedCurrent["key"] = "value"
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}

	cfg.SmartReset()
	expectedOld = make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedCurrent = make(map[string]string)
	expectedCurrent["key"] = "value"
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}

}

func TestSmartReset_settemporarilymultiple(t *testing.T) {
	cfg := WithSetTemporarilyAge(2)
	cfg.SetTemporarily("key", "value")
	expectedCurrent := make(map[string]string)
	expectedCurrent["key"] = "value"
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}
	expectedOld := make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedSetTempAgeTracker := make(map[string]int)
	expectedSetTempAgeTracker["key"] = 0
	if !reflect.DeepEqual(cfg.setTemporarilyAgeTracker, expectedSetTempAgeTracker) {
		t.Errorf(
			"Expected %+q but got %+q",
			expectedSetTempAgeTracker,
			cfg.setTemporarilyAgeTracker,
		)
	}

	cfg.SmartReset()
	expectedCurrent = make(map[string]string)
	expectedCurrent["key"] = "value"
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}
	expectedOld = make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedSetTempAgeTracker = make(map[string]int)
	expectedSetTempAgeTracker["key"] = 1
	if !reflect.DeepEqual(cfg.setTemporarilyAgeTracker, expectedSetTempAgeTracker) {
		t.Errorf(
			"Expected %+q but got %+q",
			expectedSetTempAgeTracker,
			cfg.setTemporarilyAgeTracker,
		)
	}

	cfg.SetTemporarily("key2", "value2")
	expectedCurrent = make(map[string]string)
	expectedCurrent["key"] = "value"
	expectedCurrent["key2"] = "value2"
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}
	expectedOld = make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedSetTempAgeTracker = make(map[string]int)
	expectedSetTempAgeTracker["key"] = 1
	expectedSetTempAgeTracker["key2"] = 0
	if !reflect.DeepEqual(cfg.setTemporarilyAgeTracker, expectedSetTempAgeTracker) {
		t.Errorf(
			"Expected %+q but got %+q",
			expectedSetTempAgeTracker,
			cfg.setTemporarilyAgeTracker,
		)
	}

	cfg.SmartReset()
	expectedCurrent = make(map[string]string)
	expectedCurrent["key"] = "value"
	expectedCurrent["key2"] = "value2"
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}
	expectedOld = make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedSetTempAgeTracker = make(map[string]int)
	expectedSetTempAgeTracker["key"] = 2
	expectedSetTempAgeTracker["key2"] = 1
	if !reflect.DeepEqual(cfg.setTemporarilyAgeTracker, expectedSetTempAgeTracker) {
		t.Errorf(
			"Expected %+q but got %+q",
			expectedSetTempAgeTracker,
			cfg.setTemporarilyAgeTracker,
		)
	}

	cfg.SmartReset()
	expectedCurrent = make(map[string]string)
	expectedCurrent["key2"] = "value2"
	if !reflect.DeepEqual(cfg.current, expectedCurrent) {
		t.Errorf("Expected %+q but got %+q", expectedCurrent, cfg.current)
	}
	expectedOld = make(map[string]string)
	if !reflect.DeepEqual(cfg.old, expectedOld) {
		t.Errorf("Expected %+q but got %+q", expectedOld, cfg.old)
	}
	expectedSetTempAgeTracker = make(map[string]int)
	expectedSetTempAgeTracker["key2"] = 2
	if !reflect.DeepEqual(cfg.setTemporarilyAgeTracker, expectedSetTempAgeTracker) {
		t.Errorf(
			"Expected %+q but got %+q",
			expectedSetTempAgeTracker,
			cfg.setTemporarilyAgeTracker,
		)
	}
}
