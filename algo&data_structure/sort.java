class Solution {
    public int[] sortArray(int[] nums) {
        quickSort(nums,0,nums.length);
        return nums;
    }
  
  //basic quick sort(0优化)
    public void quickSort(int[] nums,int l,int r){
        if(l >= r){
            return;
        }
        int n = partition(nums,l,r);
        quickSort(nums,l,n);
        quickSort(nums,n+1,r);
        return;
    }
    public int partition(int[] nums,int l,int r){//[)
        int pivot = nums[l];
        int lt = l;
        for(int i = l;i < r;i++){
            if(nums[i] < pivot){
                lt++;
                swap(nums,i,lt);
            }
        }
        swap(nums,l,lt);
        return lt;
    }
    public void swap(int[] nums,int i,int j){
        int t = nums[i];
        nums[i] = nums[j];
        nums[j] = t;
    }
}
