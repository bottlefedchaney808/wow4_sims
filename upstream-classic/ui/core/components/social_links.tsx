import tippy from 'tippy.js';

import { Component } from './component';

export class SocialLinks extends Component {
	static buildDiscordLink(): Element {
		const anchor = (
			<a href="https://discord.gg/p3DgvmnDCS" target="_blank" className="discord-link link-alt" dataset={{ tippyContent: 'Join us on Discord' }}>
				<i className="fab fa-discord fa-lg" />
			</a>
		);
		tippy(anchor);
		return anchor;
	}
}
