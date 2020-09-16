import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:mod_write/mod_write.dart';

void main() => runApp(ModularApp(
      module: AppModule(),
    ));

class AppModule extends MainModule {
  // here will be any class you want to inject into your project (eg bloc, dependency)
  @override
  List<Bind> get binds => [];

  // here will be the routes of your module
  @override
  List<ModularRouter> get routers => [
        ModularRouter(
          "/",
          module: ModWriterModule("/"),
        )
      ];

// add your main widget here
  @override
  Widget get bootstrap => AppWidget();
}

class AppWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      navigatorKey: Modular.navigatorKey,
      debugShowCheckedModeBanner: false,
      // set your initial route
      initialRoute: "/",
      // add Modular to manage the routing system
      onGenerateRoute: Modular.generateRoute,
    );
  }
}
