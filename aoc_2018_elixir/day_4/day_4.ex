defmodule Day4 do
  def solve_part_1() do
    File.stream!("input.txt")
    |> format_input()
    |> get_guard_sleep_tabel()
    |> guard_with_most_sleep()
    |> get_most_occurred_sleep_time()
    |> give_answer_part_1()
  end

  def solve_part_2() do
    File.stream!("input.txt")
    |> format_input()
    |> get_guard_sleep_tabel()
    |> Enum.map(&get_most_occurred_sleep_time/1)
    |> Enum.max_by(fn {_id, {_minute, sleep_time}} -> sleep_time end)
    |> give_answer_part_2()
  end

  defp give_answer_part_2({id, {minute, _sleep_time}}), do: id * minute

  defp give_answer_part_1({id, {slept_most_at, _sleep_time}}), do: id * slept_most_at

  defp guard_with_most_sleep(sleep_table) do
    Enum.max_by(sleep_table, fn {_id, time_table} ->
      Enum.reduce(time_table, 0, fn {fell_asleep, woke_up}, total_sleep ->
        total_sleep + woke_up - fell_asleep
      end)
    end)
  end

  defp get_most_occurred_sleep_time({id, sleep_schedule}) do
    most_sleep_occurred_at =
      Enum.reduce(sleep_schedule, %{}, fn {fell_asleep, woke_up}, sleep_time_occurance ->
        Enum.reduce(fell_asleep..(woke_up - 1), sleep_time_occurance, fn sleep_time,
                                                                         sleep_time_occurance ->
          Map.update(sleep_time_occurance, sleep_time, 1, fn times -> times + 1 end)
        end)
      end)
      |> Enum.max_by(fn {_time, occurances} ->
        occurances
      end)

    {id, most_sleep_occurred_at}
  end

  def get_guard_sleep_tabel(sorted_guard_schedule) do
    Enum.reduce(sorted_guard_schedule, {_guard = 0, _fell_asleep = 0, %{}}, fn
      {_date, "Guard #" <> action}, {_guard, fell_asleep, sleep_tracker} ->
        new_guard = get_guard_id(action)

        {new_guard, fell_asleep, sleep_tracker}

      {date, "falls asleep"}, {guard, _fell_asleep, sleep_tracker} ->
        fell_asleep_at = get_time(date)

        {guard, fell_asleep_at, sleep_tracker}

      {date, "wakes up"}, {guard, fell_asleep, sleep_tracker} ->
        wakes_up_at = get_time(date)
        tracked_sleep = {fell_asleep, wakes_up_at}

        updated_sleep_tracker =
          Map.update(sleep_tracker, guard, [tracked_sleep], fn sleep_list ->
            [tracked_sleep | sleep_list]
          end)

        {guard, fell_asleep, updated_sleep_tracker}
    end)
    |> elem(2)
  end

  defp get_guard_id(action) do
    String.split(action) |> hd |> String.to_integer()
  end

  defp get_time(date) do
    String.slice(date, -2, String.length(date)) |> String.to_integer()
  end

  defp format_input(input) do
    input
    |> Stream.map(&to_date_string_and_action/1)
    |> Enum.sort(fn {date1, _}, {date2, _} -> date1 < date2 end)
  end

  defp to_date_string_and_action(line) do
    {date_string, action_string} =
      String.replace(line, "\n", "")
      |> String.split_at(18)

    {to_date(date_string), String.trim(action_string)}
  end

  defp to_date(date_string) do
    date_string
    |> String.slice(1, 16)

    # Uncomment if dates are needed instead of string
    # |> (&(&1 <> ":00")).()
    # |> NaiveDateTime.from_iso8601()
    # |> elem(1)
  end
end
