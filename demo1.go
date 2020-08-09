//���������������İ�
package main

//����fmt����ʹ�����
import "fmt"

//����main������ע��Go�еĴ����ŷ�ʽֻ֧����һ����ʽ
func main() {
	//Println�����е�����js��console.log����
	fmt.Println("123456")              //123456
	fmt.Println("hello", "world", "!") //hello world !

	//��ε�������Ķ����ʽ
	//ʹ��%������ָ�v�ķ�ʽ���м�����־ͱ�ʾ���ռλ���ĳ��ȣ����ĵĻ���ÿ���ְ������ַ��㣩��������ʾ���Ҷ��룬������ʾ�������
	fmt.Printf("%-15v %6v\n", "abcdefghijklmno", "123")
	fmt.Printf("%-15v %6v\n", "abcdefg", "123456")
	fmt.Printf("%-15v %6v\n", "һ�����������߰˾�", "123456")

	//const����������var����������ûʲô�ر�֮��
	const width = 10
	var height = 5
	var distance, speed = 5600000, 10080

	fmt.Println("area = ", width*height)
	fmt.Println("time = ", distance/speed)

	fmt.Println(add(10, 2))             //12
	fmt.Println(reverse("a", "b", "c")) //c b a
	fmt.Println(split(14))              //6 8
}

//�����Ķ��巽ʽ
//����ֵΪһ�����������������ͬ�����ʡ�԰���������������
func add(a, b int) int {
	return a + b
}

//ţ�Ƶĵط����ˣ�֧�ֶ������ֵ
func reverse(a, b, c string) (string, string, string) {
	return c, b, a
}

//��һ��ţ�Ƶ�д������������ֵ��û�в����� return ��䷵���������ķ���ֵ��
//ֱ�ӷ������Ӧ���������¶̺����С������Ӱ�����Ŀɶ��ԡ�
//��������в���y��ֵ����y����Ĭ��ֵ��int����0��
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
