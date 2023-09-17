import 'dart:convert';

import 'package:frontend/src/search_feature/search_view.dart';
import 'package:http/http.dart';

class ShallowRecipe {
  final int id;
  final String name;

  const ShallowRecipe({required this.id, required this.name});

  factory ShallowRecipe.fromJson(Map<String, dynamic> json) {
    return ShallowRecipe(
      id: json['id'],
      name: json['name'],
    );
  }
}

Future<List<ShallowRecipe>> fetchRecipes() async {
  final response =
      await get(Uri.parse("http://localhost:8080/familys/1/recipes"));

  if (response.statusCode == 200) {
    List<dynamic> data = jsonDecode(response.body)["data"];
    List<ShallowRecipe> list =
        data.map((data) => ShallowRecipe.fromJson(data)).toList();
    return list;
  } else {
    throw Exception("Failed to load recipes");
  }
}

Future<int> deleteRecipe(int id) async {
  final response = await delete(Uri.parse("http://localhost:8080/recipes/$id"));
  return response.statusCode;
}

class FullRecipe {
  final int id;
  final String name;
  final double defaultPortions;
  final String defaultMeal;
  final List<RecipeIngredient> recipeIngredients;

  const FullRecipe(
      {required this.id,
      required this.name,
      required this.defaultPortions,
      required this.defaultMeal,
      required this.recipeIngredients});
// required this.recipeIngredients
  factory FullRecipe.fromJson(Map<String, dynamic> json) {
    return FullRecipe(
        id: json["id"],
        name: json["name"],
        defaultPortions: (json["defaultPortions"] as num).toDouble(),
        defaultMeal: json["defaultMeal"],
        recipeIngredients: json["recipeIngredients"]
            .map<RecipeIngredient>(
                (ingredient) => RecipeIngredient.fromJson(ingredient))
            .toList());
  }
}

class RecipeIngredient {
  final int id;
  final String name;
  final double amountPerPortion;
  final String unit;
  final String market;
  final bool isBio;

  const RecipeIngredient(
      {required this.id,
      required this.name,
      required this.amountPerPortion,
      required this.unit,
      required this.market,
      required this.isBio});

  factory RecipeIngredient.fromJson(Map<String, dynamic> json) {
    return RecipeIngredient(
        id: json["id"],
        name: json["name"],
        amountPerPortion: (json["amountPerPortion"] as num).toDouble(),
        unit: json["unit"],
        market: json["market"],
        isBio: json["isBio"]);
  }
}

Future<FullRecipe> fetchRecipe(int id) async {
  final response = await get(Uri.parse("http://localhost:8080/recipes/$id"));

  if (response.statusCode == 200) {
    final recipe = FullRecipe.fromJson(jsonDecode(response.body)["data"]);
    return recipe;
  } else {
    throw Exception("Failed to load recipe");
  }
}

Future<void> deleteRecipeIngredient(int id) async {
  final response = await delete(
      Uri.parse("http://localhost:8080/recipes/recipeingredient/$id"));
  if (response.statusCode == 200) {
    return;
  } else {
    throw Exception("Failed to delete recipe");
  }
}

class RecipeId {
  final int id;

  const RecipeId({
    required this.id,
  });
// required this.recipeIngredients
  factory RecipeId.fromJson(Map<String, dynamic> json) {
    return RecipeId(id: json["id"]);
  }
}

Future<int> postRecipe(String name, double defaultPortions, String meal) async {
  final response = await post(
      Uri.parse("http://localhost:8080/familys/1/recipes"),
      headers: <String, String>{
        "Content-Type": "application/json; charset=UTF-8"
      },
      body: jsonEncode(<String, dynamic>{
        "name": name,
        "defaultPortions": defaultPortions,
        "defaultMeal": meal,
      }));

  if (response.statusCode == 200) {
    return jsonDecode(response.body)["data"]["id"];
  } else {
    throw Exception("Failed to post recipe");
  }
}

Future<void> postRecipeIngredient(
    int recipeId, PostRecipeIngredientBody body) async {
  final response = await post(
      Uri.parse("http://localhost:8080/recipes/$recipeId/recipeingredient"),
      headers: <String, String>{
        "Content-Type": "application/json; charset=UTF-8"
      },
      body: jsonEncode(<String, dynamic>{
        "ingredientId": body.ingredientId,
        "amountPerPortion": body.amountPerPortion,
        "unit": body.unit,
        "market": body.market,
        "isBio": body.isBio
      }));

  if (response.statusCode == 200) {
    return;
  } else {
    throw Exception("Failed to post recipe");
  }
}