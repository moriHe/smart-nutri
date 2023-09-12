import 'package:flutter/material.dart';
import 'package:frontend/src/api/recipes.dart';
import 'package:frontend/src/my_recipes_feature/meals.dart';
import 'package:frontend/src/my_recipes_feature/recipe_details_view.dart';

List<String> listOfValue = ['1', '2', '3', '4', '5'];

class _MyRecipesState extends State<MyRecipesListView> {
  late Future<List<ShallowRecipe>> futureRecipes;
  String mealsDropDownValue = mealsKeys.first;
  final recipeNameController = TextEditingController();
  final defaultPortionsController = TextEditingController();

  @override
  void initState() {
    super.initState();
    futureRecipes = fetchRecipes();
  }

  @override
  void dispose() {
    recipeNameController.dispose();
    defaultPortionsController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text("Meine Rezepte")),
      floatingActionButtonLocation: FloatingActionButtonLocation.endTop,
      floatingActionButton: FloatingActionButton(
        onPressed: () async {
          showDialog(
              context: context,
              builder: (BuildContext context) {
                return AlertDialog(
                  scrollable: true,
                  title: const Text('Login'),
                  content: Padding(
                    padding: const EdgeInsets.all(8.0),
                    child: Form(
                      child: Column(
                        children: <Widget>[
                          TextFormField(
                            controller: recipeNameController,
                            decoration: const InputDecoration(
                              labelText: 'Name',
                              icon: Icon(Icons.article_outlined),
                            ),
                          ),
                          TextFormField(
                            controller: defaultPortionsController,
                            keyboardType: TextInputType.number,
                            decoration: const InputDecoration(
                              labelText: 'Portionen',
                              icon: Icon(Icons.brunch_dining_outlined),
                            ),
                          ),
                          const SizedBox(
                            height: 25.0,
                          ),
                          Row(
                            children: [
                              const Icon(
                                Icons.wb_sunny_outlined,
                                color: Colors.grey,
                              ),
                              const SizedBox(width: 16.0),
                              DropdownMenu<String>(
                                label: const Text("Mahlzeit"),
                                initialSelection: meals[mealsKeys.first],
                                onSelected: (String? value) {
                                  // This is called when the user selects an item.
                                  setState(() {
                                    mealsDropDownValue = value!;
                                  });
                                },
                                dropdownMenuEntries: mealsKeys
                                    .map<DropdownMenuEntry<String>>(
                                        (String value) {
                                  return DropdownMenuEntry<String>(
                                      value: meals[value]!,
                                      label: meals[value]!);
                                }).toList(),
                              ),
                            ],
                          ),
                        ],
                      ),
                    ),
                  ),
                  actions: [
                    TextButton(
                        style: TextButton.styleFrom(
                            backgroundColor: Colors.blue,
                            foregroundColor: Colors.white,
                            padding: const EdgeInsets.all(16.0),
                            textStyle: const TextStyle(fontSize: 20)),
                        onPressed: () async {
                          double? defaultPortions =
                              double.tryParse(defaultPortionsController.text);
                          if (defaultPortions == null) {
                            return;
                          }
                          await postRecipe(
                              recipeNameController.text,
                              defaultPortions,
                              meals.keys.firstWhere(
                                  (key) => meals[key] == mealsDropDownValue,
                                  orElse: () => "NONE"));
                          if (context.mounted) {
                            setState(() {
                              futureRecipes = fetchRecipes();
                            });
                            Navigator.pop(context);
                          }
                        },
                        child: const Text("Hinzufügen"))
                  ],
                );
              }).then((value) {
            mealsDropDownValue = mealsKeys.first;
            recipeNameController.text = "";
            defaultPortionsController.text = "";
          });
        },
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
