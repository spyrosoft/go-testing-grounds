#define character_array_length 80

char example[character_array_length];

char* get_char_array(void) {
	return example;
}

int get_char_array_length(void) {
	return character_array_length;
}