package asciiart

import (
	"testing"
)

func TestGetAsciiType(t *testing.T) {
	t.Run("!", testGetAsciiFunc(33, '!'))
	t.Run("\n", testGetAsciiFunc(10, '\n'))
}

func testGetAsciiFunc(n byte, expected rune) func(*testing.T) {
	return func(t *testing.T) {
		actual := GetASCII(n)
		if actual != expected {
			t.Errorf("Expected %v but instead got %v!", expected, actual)
		}
	}
}

func TestBlockelems(t *testing.T) {
	expected := "▀ ▁ ▂ ▃ ▄ ▅ ▆ ▇ █ ▉ ▊ ▋ ▌ ▍ ▎ ▏ ▐ ░ ▒ ▓ ▔ ▕ ▖ ▗ ▘ ▙ ▚ ▛ ▜ ▝ ▞ ▟"
	actual := Blockelems()

	if actual != expected {
		t.Errorf("Expected %s but instead got %s!", expected, actual)
	}
}

func TestRenderPNMBW(t *testing.T) {
	var img1 = "P2 24 7 7 0 0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0 0  1  1  1  1  0  0  2  2  2  2  0  0 3 3 3 3  0  0 4 4 4 4  0 0  1  0  0  0  0  0  2  0  0  0  0  0 3  0  0  0  0  0 4  0  0 4  0 0  1  1  1  0  0  0  2    2  2  0  0  0 3 3 3  0  0  0 4 4 4 4  0 0  1  0  0  0  0  0  2  0  0  0  0  0 3  0  0  0  0  0 4  0  0  0  0 0  1  0  0  0  0  0  2  2  2  2  0  0 3 3 3 3  0  0 4  0  0  0  0 0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0 "
	var result1 = `
                        
 ░░░░  ▒▒▒▒  ▓▓▓▓  ████ 
 ░     ▒     ▓     █  █ 
 ░░░   ▒▒▒   ▓▓▓   ████ 
 ░     ▒     ▓     █    
 ░     ▒▒▒▒  ▓▓▓▓  █    
                        
`

	var img2 = `P1          #Bitmap
# Beispiel für das Bild des Buchstabens "U"
6 9        #Breite des Bildes, Leerstelle, Höhe des Bildes
1 0 0 0 1 0
1 0 0 0 1 0
1 0 0 0 1 0
1 0 0 0 1 0
1 0 0 0 1 0
1 0 0 0 1 0
1 0 0 0 1 0
0 1 1 1 0 0
0 0 0 0 0 0`
	var result2 = `
█   █ 
█   █ 
█   █ 
█   █ 
█   █ 
█   █ 
█   █ 
 ███  
      
`

	var img3 = `P2
# Das Wort "FEEP" in verschiedenen Graustufen 
# (Beispiel von der Netpbm-Man-Page)
24 7 
15 
0 0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0 
0  3  3  3  3  0  0  7  7  7  7  0  0 11 11 11 11  0  0 15 15 15 15  0
0  3  0  0  0  0  0  7  0  0  0  0  0 11  0  0  0  0  0 15  0  0 15  0
0  3  3  3  0  0  0  7    7  7  0  0  0 11 11 11  0  0  0 15 15 15 15  0
0  3  0  0  0  0  0  7  0  0  0  0  0 11  0  0  0  0  0 15  0  0  0  0
0  3  0  0  0  0  0  7  7  7  7  0  0 11 11 11 11  0  0 15  0  0  0  0
0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0
`

	t.Run("feep.pnm", testPNBWFunc(img1, result1))
	t.Run("U.pnm", testPNBWFunc(img2, result2))
	t.Run("feep2.pnm", testPNBWFunc(img3, result1))
}

func testPNBWFunc(img string, expected string) func(*testing.T) {
	return func(t *testing.T) {
		actual := RenderPNMBW(img)
		if actual != expected {
			t.Errorf("Expected %s but instead got %s!", expected, actual)
		}
	}
}

func TestRenderASC(t *testing.T) {
	var img1 = `

	������������������������������������������������������������������߲
	�       ��                                                         �
	�                                                                  �
	�                           �               �         drm          �
������  ������ �  ������ ������ ���� ������ ��� ���������      ���������
�۲����   ������  �۲��� �۲���  ����� ��۲�� ����� ��������� ���� ���������
�������  �  �����  ������ ������������   ����� ������ �  ����������� �  ��۲��
�۲���  � ���۲��۲��� ��۲��������� � ޲���� �۲��  �� ����� ����  �  ������  � ��������� ����������� �����   ۲���    ����� ������   �۲���������   ��۲���  ߲�  ���������      ��     ��� ������   ���   � �۲���  ������  ��߲��������
	�  Vihino         ������      �����                           ��   �
	�                                                                  �
	�               xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx               �
	�               xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx               �
	�               xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx               �
	�� ��           xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx         �     �
	�               xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx               �
	�               xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx               �
	�                                                                  �
	��������������������������������������������������������������������
SAUCE00
`

	var result1 = ""

	t.Run("SandS.asc", testASCFunc(img1, result1))

}

func testASCFunc(img string, expected string) func(*testing.T) {
	return func(t *testing.T) {
		actual := RenderASC(img)
		if actual != expected {
			t.Errorf("Expected %s but instead got %s!", expected, actual)
		}
	}
}
