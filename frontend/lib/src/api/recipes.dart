import 'dart:convert';

import 'package:http/http.dart';

class Recipe {
  final int id;
  final String name;

  const Recipe({required this.id, required this.name});

  factory Recipe.fromJson(Map<String, dynamic> json) {
    return Recipe(
      id: json['id'],
      name: json['name'],
    );
  }
}

Future<Recipe> fetchRecipes() async {
  final response =
      await get(Uri.parse("http://localhost:8080/familys/1/recipes"));

  if (response.statusCode == 200) {
    Map<String, dynamic> body = json.decode(response.body);
    List<dynamic> list = body["data"];
    Recipe recipe = Recipe.fromJson(list[0]);
    return recipe;
  } else {
    throw Exception("Failed to load recipes");
  }
}
