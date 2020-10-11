def next_line
  gets.chomp.split(' ').map(&:to_i)
end


if __FILE__ == $0
  n, d = next_line
  distance = d**2

  answer = n.times.reduce(0) do |cnt, _|
    x, y = next_line
    x**2 + y**2 <= distance ? cnt + 1 : cnt
  end

  puts answer
end
