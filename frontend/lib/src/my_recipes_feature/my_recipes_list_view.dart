import 'package:flutter/material.dart';
import 'package:frontend/src/api/recipes.dart';

class _MyAppState extends State<MyRecipesListView> {
  late Future<Recipe> futureRecipes;

  @override
  void initState() {
    super.initState();
    futureRecipes = fetchRecipes();
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Fetch Data Example',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: Scaffold(
        appBar: AppBar(
          title: const Text('Fetch Data Example'),
        ),
        body: Center(
          child: FutureBuilder<Recipe>(
            future: futureRecipes,
            builder: (context, snapshot) {
              if (snapshot.hasData) {
                return Text(snapshot.data!.name);
              } else if (snapshot.hasError) {
                return Text('${snapshot.error}');
              }

              // By default, show a loading spinner.
              return const CircularProgressIndicator();
            },
          ),
        ),
      ),
    );
  }
}

class MyRecipesListView extends StatefulWidget {
  const MyRecipesListView({super.key});

  @override
  State<MyRecipesListView> createState() => _MyAppState();

  static const routeName = "/my-recipes";

  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(title: const Text("Meine Rezepte")),
        body: const Center(
          child: Text("Meine Rezepte Body"),
        ));
  }
}
