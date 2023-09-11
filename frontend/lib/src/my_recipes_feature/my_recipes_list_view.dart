import 'package:flutter/material.dart';
import 'package:frontend/src/api/recipes.dart';
import 'package:frontend/src/my_recipes_feature/recipe_details_view.dart';

class _MyRecipesState extends State<MyRecipesListView> {
  late Future<List<ShallowRecipe>> futureRecipes;

  @override
  void initState() {
    super.initState();
    futureRecipes = fetchRecipes();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text("Meine Rezepte")),
      floatingActionButtonLocation: FloatingActionButtonLocation.endTop,
      floatingActionButton: FloatingActionButton(
        onPressed: () {},
        backgroundColor: Colors.green,
        child: const Icon(Icons.add),
      ),
      body: Center(
        child: FutureBuilder<List<ShallowRecipe>>(
          future: futureRecipes,
          builder: (context, snapshot) {
            if (snapshot.hasData) {
              return Center(
                  child: ListView.builder(
                      padding: const EdgeInsets.only(top: 50.0),
                      itemCount: snapshot.data!.length,
                      itemBuilder: (BuildContext context, int index) {
                        return Center(
                            child: GestureDetector(
                                onTap: () => Navigator.pushNamed(
                                    context, RecipeDetailsView.routeName,
                                    arguments: RecipeDetailsViewArguments(
                                        snapshot.data![index].id)),
                                child: Card(
                                    margin: const EdgeInsets.only(
                                        right: 20.0,
                                        left: 20.0,
                                        top: 10.0,
                                        bottom: 10.0),
                                    child: ListTile(
                                        title: Center(
                                      child: Text(snapshot.data![index].name),
                                    )))));
                      }));
            } else if (snapshot.hasError) {
              return Text('${snapshot.error}');
            }

            // By default, show a loading spinner.
            return const CircularProgressIndicator();
          },
        ),
      ),
    );
  }
}

class MyRecipesListView extends StatefulWidget {
  const MyRecipesListView({super.key});

  @override
  State<MyRecipesListView> createState() => _MyRecipesState();

  static const routeName = "/my-recipes";
}
