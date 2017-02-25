package main

import (
	"math/rand"
	"net/http"
)

var oneliners = []string{
	"If i had a dollar for every girl that found me unattractive, they would eventually find me attractive.",
	"Feeling pretty proud of myself. The Sesame Street puzzle I bought said 3-5 years, but I finished it in 18 months.",
	"Today a man knocked on my door and asked for a small donation towards the local swimming pool. I gave him a glass of water.",
	"Strong people don't put others down. They lift them up and slam them on the ground for maximum damage.",
	"Life is all about perspective. The sinking of the Titanic was a miracle to the lobsters in the ship's kitchen.",
	"You know you're ugly when it comes to a group picture and they hand you the camera.",
	"Just read that 4,153,237 people got married last year, not to cause any trouble but shouldn't that be an even number?",
	"Life is like toilet paper, you're either on a roll or taking shit from some asshole.",
	"I want to die peacefully in my sleep, like my grandfather. Not screaming and yelling like the passengers in his car.",
	"When wearing a bikini, women reveal 90 % of their body... men are so polite they only look at the covered parts.",
	"I used to think I was indecisive, but now I'm not too sure.",
	"You know that tingly little feeling you get when you like someone? That's your common sense leaving your body.",
	"When my boss asked me who is the stupid one, me or him? I told him everyone knows he doesn't hire stupid people.",
	"Team work is important; it helps to put the blame on someone else.",
	"I'm great at multitasking. I can waste time, be unproductive, and procrastinate all at once.",
	"Apparently I snore so loudly that it scares everyone in the car I'm driving.",
	"Before I criticize a man, I like to walk a mile in his shoes. That way, when I do criticize him, I'm a mile away and I have his shoes.",
	"Intelligence is like an underwear. It is important that you have it, but not necessary that you show it off.",
	"If you think nobody cares whether you're alive, try missing a couple of payments.",
	"I can totally keep secrets. It's the people I tell them to that can't.",
	"Take my advice — I'm not using it.",
	"I think my neighbor is stalking me as she's been googling my name on her computer. I saw it through my telescope last night.",
	"My therapist says I have a preoccupation with vengeance. We'll see about that.",
	"When an employment application asks who is to be notified in case of emergency, I always write, \"A very good doctor\".",
	"Did you know that dolphins are so smart that within a few weeks of captivity, they can train people to stand on the very edge of the pool and throw them fish?",
	"I changed my password to \"incorrect\". So whenever I forget what it is the computer will say \"Your password is incorrect\".",
	"I started out with nothing, and I still have most of it.",
	"There are two rules for success: 1) Don't tell all you know",
	"Artificial intelligence is no match for natural stupidity.",
	"I bought a vacuum cleaner six months ago and so far all it's been doing is gathering dust.",
	"The 50-50-90 rule: Anytime you have a 50-50 chance of getting something right, there's a 90% probability you'll get it wrong.",
	"Entered what I ate today into my new fitness app and it just sent an ambulance to my house.",
	"Good health is merely the slowest possible rate at which one can die.",
	"Never test the depth of the water with both feet.",
	"When I call a family meeting I turn off the house wifi and wait for them all to come running.",
	"A conclusion is the part where you got tired of thinking.",
	"Thanks for explaining the word \"many\" to me, it means a lot.",
	"A computer once beat me at chess, but it was no match for me at kick boxing.",
	"I hate when I am about to hug someone really sexy and my face hits the mirror.",
	"A clean house is the sign of a broken computer.",
	"If you're not supposed to eat at night, why is there a light bulb in the refrigerator?",
	"Improve your memory by doing unforgettable things.",
	"People used to laugh at me when I would say \"I want to be a comedian\", well nobody's laughing now.",
	"When I see ads on TV with smiling, happy housewives using a new cleaning product, the only thing I want to buy are the meds they must be on.",
	"Dear alcohol, We had a deal where you would make me funnier, smarter, and a better dancer... I saw the video... we need to talk.",
	"Waking up this morning was an eye-opening experience.",
	"There are three kinds of people: Those who can count and those who can't.",
	"I hate people who use big words just to make themselves look perspicacious.",
	"Keep the dream alive: Hit the snooze button.",
	"I got caught in police speed trap yesterday. The officer walked up to my car and said \"I've been waiting all day for you\" Well I said. I got here as fast as I could.",
	"I am a nobody, nobody is perfect, therefore I am perfect.",
	"I'm really good at stuff until people watch me do that stuff.",
	"Nothing ruins a Friday more than an understanding that today is Tuesday.",
	"A straight face and a sincere-sounding \"Huh?\" have gotten me out of more trouble than I can remember.",
	"For maximum attention, nothing beats a good mistake.",
	"I didn't say it was your fault, I said I was blaming you.",
	"If time is money are ATMs time machines?",
	"A fine is a tax for doing wrong. A tax is a fine for doing well.",
	"I think it's wrong that only one company makes the game Monopoly.",
	"If procrastionation was an Olympic sport, I'd compete in it later.",
}

// OneLinerHandlerFunc is the oneliner API. It will write a random one-line joke.
func OneLinerHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	index := rand.Int() % len(oneliners)
	w.Write([]byte(oneliners[index]))
}
