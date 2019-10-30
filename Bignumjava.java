public class Bignumjava {
	public static void main(String args[]){
		long n=1000000000, sum=0;
		System.out.print("n = ");
		System.out.println(n);
		for (int i=0; i<n; i++) {
			sum+=i;
		}
		System.out.print("sum = ");
		System.out.println(sum);
	}
}