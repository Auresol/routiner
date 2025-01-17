import 'package:flutter/material.dart';
import 'package:routiner/src/repo/task_repo.dart';

import 'feature/home_page.dart';
import 'feature/edit_page.dart';
import 'feature/setting_page.dart';

import 'package:routiner/src/repo/routine_repo.dart';
import 'package:provider/provider.dart';

final pages = [
  [HomePage(), Icons.home, "home"], 
  [const EditPage(), Icons.edit, "edit"], 
  [const SettingPage(), Icons.settings, "settings"]
];

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MultiProvider(
      providers: [
        ChangeNotifierProvider(create: (context) => RoutineProvider()),
        ChangeNotifierProvider(create: (context) => TaskProvider()), 
      ],
      child: MaterialApp(
        title: 'Flutter Demo',
        theme: ThemeData(
          colorScheme: ColorScheme.fromSeed(seedColor: Colors.lightGreen),
          useMaterial3: true,
        ),
        home: const DesktopPage(title: 'Routiner'),
      ),
    );
  }
}

class DesktopPage extends StatefulWidget {
  const DesktopPage({super.key, required this.title});

  final String title;

  @override
  State<DesktopPage> createState() => _MyDesktopPageState();
}

class _MyDesktopPageState extends State<DesktopPage> {

  var selectedIndex = 0;

  @override
  Widget build(BuildContext context) {

    Widget page;
    page = pages[selectedIndex][0] as Widget;


    return LayoutBuilder(
      builder: (context, constraints) {
        return Scaffold(
          body: Row(
            children: [
              SafeArea(
                child: NavigationRail(
                  backgroundColor: Theme.of(context).primaryColorLight,
                  extended: constraints.maxWidth >= 600,
                  destinations: [
                    for(var item in pages)
                      NavigationRailDestination(icon: Icon(item[1] as IconData), label: Text(item[2].toString()),)
                  ],
                  selectedIndex: selectedIndex,
                  onDestinationSelected: (value) {
                    
                    // ↓ Replace print with this.
                    setState(() {
                      selectedIndex = value;
                    });
        
                  },
                ),
              ),
              Expanded(
                child: Container(
                  //color: Theme.of(context).colorScheme.primaryContainer,
                  child: page,
                ),
              ),
            ],
          ),
        );
      }
    );
  }
}