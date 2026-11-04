import { Component } from '@angular/core';
import { ChatInput } from './chat-input/chat-input';
import { MessageModel } from '../../shared/models';
import { Message } from '../../shared/message/message';


@Component({
    imports: [ChatInput, Message, Message],
    selector: 'server-chat',
    styleUrl: './server-chat.css',
    templateUrl: './server-chat.html',
})
export class ServerChat {
    chatMessages: MessageModel[] = [
        {
            user: { name: 'WASSUP' },
            msg: 'Всем привет, кто тут?',
        },
        {
            user: { name: 'HoJLter' },
            msg: 'Я тут, ковыряю интерфейс.',
        },
        {
            user: { name: 'MoonWalker' },
            msg: 'О, уже можно сообщения отправлять?',
        },
        {
            user: { name: 'HoJLter' },
            msg: 'Пока заглушки, но выглядит почти как настоящее общение.',
        },
        {
            user: { name: 'WASSUP' },
            msg: 'Получается, и я заглушка?',
        },
        {
            user: { name: 'MoonWalker' },
            msg: 'У меня для тебя плохие новости…',
        },
        {
            user: { name: 'alex_dev' },
            msg: 'Закиньте длинное сообщение, надо проверить переносы и высоту блока. Например, вот такое: если текст занимает несколько строк, аватарка должна оставаться сверху, а время сообщения не должно уезжать за границы контейнера.',
        },
        {
            user: { name: 'WASSUP' },
            msg: 'Ок',
        },
        {
            user: { name: 'night_owl' },
            msg: 'А тёмная тема уже есть? 🌙',
        },
        {
            user: { name: 'HoJLter' },
            msg: 'У нас пока только тёмная.',
        },
        {
            user: { name: 'alex_dev' },
            msg: 'ПроверкаДлинногоСловаБезПробеловЧтобыНичегоНеВылезалоЗаГраницыКонтейнера',
        },
        {
            user: { name: 'MoonWalker' },
            msg: 'Первая строка\nВторая строка\nТретья строка',
        },
        {
            user: { name: 'night_owl' },
            msg: 'Выглядит неплохо 👍',
        },
        {
            user: { name: 'WASSUP' },
            msg: 'Ладно, пишите, когда перестану быть заглушкой.',
        },
    ];
}
