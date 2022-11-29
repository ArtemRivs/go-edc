// объявление типа
type Special int 
// объявление метода
func (m Special) GetString() string{
      return fmt.Sprintf("Special: %d", m)      
}
// вызов метода
var m Special = 5
var s string = m.GetString()
