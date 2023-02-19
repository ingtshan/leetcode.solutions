# @param {Integer[]} nums
# @param {Integer} target
# @return {Integer[]}
def two_sum(nums, target)
  hash = {}
  nums.each_with_index do |n, i|
    her = target - n
    return [hash[her], i] unless hash[her].nil?
    hash[n] = i
  end
end
