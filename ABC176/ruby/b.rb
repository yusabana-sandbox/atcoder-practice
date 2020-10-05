input = gets.chomp
 
sum = input.split("").reduce(0) { |r, a| r + a.to_i }
 
puts(sum % 9 == 0 ? "Yes" : "No")
