import 'dart:convert';

import 'package:flutter/cupertino.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/widgets.dart';
import 'package:flutter/services.dart';
import 'package:sys_core/sys_core.dart';
import './translations.dart';

class KanbanLocalizations extends Translations {
  final Locale locale;
  Map<String, String> _localizedStrings;

  KanbanLocalizations(this.locale);

  Future<bool> load() async {
    String jsonString = await rootBundle.loadString(
        'packages/mod_kanban/i18n/lang_${locale.languageCode}.json');

    Map<String, dynamic> jsonMap = Map.from(json.decode(jsonString))
      ..removeWhere((key, value) => key[0] == '@');

    this._localizedStrings = jsonMap.map((key, value) {
      return MapEntry(key, value.toString());
    });

    return true;
  }

  String translate(String key) {
    return _localizedStrings[key];
  }

  static KanbanLocalizations of(BuildContext context) {
    return Localizations.of<KanbanLocalizations>(
        context, KanbanLocalizations);
  }
}

class KanbanLocalizationsDelegate
    extends LocalizationsDelegate<KanbanLocalizations> {
  final Locale overriddenLocale;

  KanbanLocalizationsDelegate(this.overriddenLocale);

  @override
  bool shouldReload(KanbanLocalizationsDelegate old) => true;

  @override
  bool isSupported(Locale locale) {
    return Languages.supportedLanguages.keys.contains(locale.languageCode);
  }

  @override
  Future<KanbanLocalizations> load(Locale locale) async {
    KanbanLocalizations localizations = new KanbanLocalizations(locale);
    await localizations.load();
    return localizations;
  }
}

class FallbackCupertinoLocalisationsDelegate
    extends LocalizationsDelegate<CupertinoLocalizations> {
  const FallbackCupertinoLocalisationsDelegate();

  @override
  bool isSupported(Locale locale) =>
      Languages.supportedLanguages.keys.contains(locale.languageCode);

  @override
  Future<CupertinoLocalizations> load(Locale locale) =>
      SynchronousFuture<_DefaultCupertinoLocalizations>(
          _DefaultCupertinoLocalizations(locale));

  @override
  bool shouldReload(FallbackCupertinoLocalisationsDelegate old) => false;
}

class _DefaultCupertinoLocalizations extends DefaultCupertinoLocalizations {
  final Locale locale;

  _DefaultCupertinoLocalizations(this.locale);
}