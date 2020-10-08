n = gets.chomp.to_i
l = gets.chomp.split(' ').map(&:to_i).sort
l_len = l.length
result = 0

(0...l_len-2).each do |i|
  (i+1...l_len-1).each do |j|
    (j+1...l_len).each do |k|
      cond1 = l[i] != l[j] && l[i] != l[k] && l[j] != l[k]
      cond2 = l[i] + l[j] > l[k]

      if cond1 && cond2
        result += 1
      end
    end
  end
end

puts result
