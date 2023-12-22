def main():
	clear_screen();
	print("Welcome to Tic Tac Toe!");
	print("------------------------------");
	menu_choice = display_menu();

	while (menu_choice != "q"):
		if (menu_choice == "s"):
			run_game();

			# Game has ended, back to main menu
			clear_screen();
		else:
			clear_screen();
			print("\nInvalid choice. Please enter a valid option.\n");
		menu_choice = display_menu();
	
	clear_screen();
	print("Thanks for playing!");

# Display menu of options and return the option value. 
def display_menu():
	print("Menu Options: ");
	print("- start game: s --------------");
	print("- quit game:  q --------------");
	print("------------------------------");
	option = input("enter: ")

	return option

# Clears screen 
def clear_screen():
	print("\n\n\n\n\n\n\n\n\n\n");
	print("------------------------------");
	
# Manages the ongoing game state
def run_game():
	board = [[1,2,3], [4,5,6], [7,8,9]]; # Board
	moves = []; # All played moves in play order starting with X
	game_is_won = False; # Check for winner
	game_is_finished = False; # Game ends with a winner or tie
	player_is_x = True; # Player 1 is X
	
	while(not game_is_finished):
		display_board(board);
		player_input = input("Choose a spot on the board: ")
		
		# If input is not valid, repeat until valid input
		while(not is_valid_move(player_input, moves)):
			display_board(board);
			player_input = input("Invalid option! Choose a spot on the board:")
		
		board = insert_move(int(player_input), player_is_x, board);
		moves.append(int(player_input));

		game_is_won = validate_win(board, int(player_input));
		
		# If game is not won, switch players
		if(not game_is_won):
			player_is_x = not player_is_x;

		game_is_finished = game_is_won or len(moves) == 9;
	
	
	display_board(board);
	if(game_is_won):
		player = "1" if player_is_x else "2"
		print("\nPlayer {} is the winner!".format(player))
	else:
		print("\nGame is a tie!")

# Displays board 		
def display_board(board = [[1,2,3], [4,5,6], [7,8,9]]):
	def draw_line():
		print("--- --- ---");
	def draw_row(row_data):
		print(" {} | {} | {} ".format(row_data[0], row_data[1], row_data[2]));

	clear_screen();
	draw_row(board[0]);
	draw_line();
	draw_row(board[1]);
	draw_line();
	draw_row(board[2]);

# Check if the move is a valid input and valid move location
# return True for valid move
def is_valid_move(new_move, moves):
	# All valid inputs should be 1-9
	if(not new_move.isnumeric()):
		return False;
	# If move was already used, then it's not valid
	if(int(new_move) in moves):
		return False;
	
	return True;

def insert_move(move, player_is_x=True, board=[[1,2,3], [4,5,6], [7,8,9]]):
	row, col = move_to_row_col(move);
	board[row][col] = "x" if player_is_x else "o";

	return board;

# Convert the move location to a row and col for use in the 2d array
# Returns (row, col)
def move_to_row_col(move):
	rows = [1,4,7];
	row = 0 if (1 <= move and move <= 3) else 1 if (4 <= move and move <= 6) else 2;
	# The column is found by subtracting the move by the first number in the row.
	# ex: 7 is in the last row and 7 is the first number in the row. 7-7 = 0 or the column position
	col = move - rows[row];
	return (row, col)

# Check if the last move resulted in a winning board. Board has all played moves.
# This only checks the last move not all possible moves.
# Returns True if the last move is in a row of 3. 
def validate_win(board, last_move):
	row, col = move_to_row_col(last_move);
	#check horizontal
	if (board[row][col] == board[row][(col+1)%3]) and (board[row][col] == board[row][(col+2)%3]):
		return True;
	#check vertical
	if (board[row][col] == board[(row+1)%3][col]) and (board[row][col] == board[(row+2)%3][col]):
		return True;
	#check left diag
	if(last_move in [3, 5, 7]):
		if (board[0][2] == board[1][1]) and (board[1][1] == board[2][0]):
			return True;
	#check right diag
	if(last_move in [1, 5, 9]):
		if (board[0][0] == board[1][1]) and (board[1][1] == board[2][2]):
			return True;

	return False



main();