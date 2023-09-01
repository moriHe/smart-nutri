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

Future<List<Recipe>> fetchRecipes() async {
  final response =
      await get(Uri.parse("http://localhost:8080/familys/1/recipes"));

  if (response.statusCode == 200) {
    List<dynamic> data = jsonDecode(response.body)["data"];
    List<Recipe> list = data.map((data) => Recipe.fromJson(data)).toList();
    return list;
  } else {
    throw Exception("Failed to load recipes");
  }
}
