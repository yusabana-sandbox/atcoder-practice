n = gets.chomp.to_i
l = gets.chomp.split(' ').map(&:to_i).sort
result = 0

(0...n-2).each do |i|
  (i+1...n-1).each do |j|
    (j+1...n).each do |k|
      cond1 = l[i] != l[j] && l[i] != l[k] && l[j] != l[k]
      cond2 = l[i] + l[j] > l[k]

      if cond1 && cond2
        result += 1
      end
    end
  end
end

puts result
