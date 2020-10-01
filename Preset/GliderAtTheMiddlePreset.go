package Preset

type GliderAtTheMiddlePreset struct {

}

func (_ *GliderAtTheMiddlePreset) InitializeField(field *[][]bool) {
	i := len(*field) / 2
	j := len((*field)[0]) / 2

	(*field)[i-1][j] = true
	(*field)[i][j+1] = true
	(*field)[i+1][j-1]=true
	(*field)[i+1][j]=true
	(*field)[i+1][j+1]=true
}
