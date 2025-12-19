package nestml

type NestML struct {
	macros map[string]NestMLMacro
	roots  []NestMLElement
}

type NestMLMacro struct {
	args []NestMLMacroArg
	body []NestMLElement
}

type NestMLMacroArg struct {
	name         string
	defaultValue *NestMLPropValue
}

type NestMLElement interface {
	String() string
}

type NestMLText string

func (t NestMLText) String() string {
	return string(t)
}

type NestMLPropKey string

type NestMLPropValue struct {
	texts  []string
	macros []string
}

type NestMLBlock struct {
	name       string
	properties map[NestMLPropKey]NestMLPropValue
	children   []NestMLElement
}

type NestMLMacroCall struct {
	name string
	args map[NestMLPropKey]NestMLPropValue
	body []NestMLElement
}
